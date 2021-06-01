package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

//模拟实现三节点选举
//改造代码成分布式选举代码，加入rpc调用

const raftConut = 3

//leader结构体
type Leader struct{
	Term int //任期
	LeaderId int //编号
}

type Raft struct {
	mu sync.Mutex
	me int //节点编号
	currentTerm int// 当前任期
	votedFor int //为那个节点投票
	state int //3个状态 0:follower 1:candidate 2:leader
	lastMessageTime int64 //发送最后一条数据的时间
	currentLeader int //当前集群领导的编号
	message chan bool
	elecCh chan bool //选举的通道
	heartBeat chan bool //心跳信号的通道
	heartbeatRe chan bool //返回心跳信号的通道
	timeout int
}

var leader = Leader{0, -1}

func main() {
	//过程，有3个节点，最初都是follwer
	//如果有candidate状态，进行投票和拉票
	//会产生leader

	//创建三个节点
	for i := 0; i < raftConut; i++ {
		//创建3个raft节点
		Make(i)
	}

	//加入rpc注册服务
	rpc.Register(new(Raft))
	rpc.HandleHTTP()
	//监听服务
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	for{}
}

func Make(me int) (rf *Raft) {
	rf = &Raft{}
	rf.me = me
	rf.votedFor = -1 //-1代表谁都不投，此时这个节点刚创建
	rf.state = 0
	rf.timeout = 0
	rf.currentLeader = -1
	rf.setTerm(0)
	rf.message = make(chan bool)
	rf.elecCh = make(chan bool)
	rf.heartBeat = make(chan bool)
	rf.heartbeatRe = make(chan bool)

	rand.Seed(time.Now().UnixNano()) //设置随机种子

	//选举协程
	go rf.election()
	//心跳检测的协程
	go rf.sendLeaderBeat()

	return
}

func (rf *Raft)setTerm(term int){
	rf.currentTerm = term
}

func (rf *Raft)election(){
	//设置标记，是否选出了leader
	var result bool
	for{
		//设置超时
		timeout := randRange(150, 300)
		rf.lastMessageTime = millisecond()
		select {
		//延时等待1ms
		case <- time.After(time.Duration(timeout) * time.Millisecond):
			fmt.Println("当前节点状态为",rf.state, rf.me)
		}
	}

	result = false
	for !result {
		//选主
		result = rf.election_one_round(&leader)
	}
}

func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

//获取当时时间，发送最后一条数据的时间
func millisecond() int64{
	return time.Now().UnixNano() / int64(time.Millisecond)
}

//实现选主的逻辑
func(rf *Raft) election_one_round(leader *Leader) bool{
	var timeout int64
	timeout = 100
	//投票数量
	var vote int
	//定义是否开始心跳信号的产生
	var triggerHeartbeat bool
	//时间
	//last := millisecond()
	//用于返回
	success := false
	//给当前节点变成candidate状态
	rf.mu.Lock()
	rf.becomeCandidate()
	rf.mu.Unlock()
	fmt.Println("start electing leader")

	for{
		//遍历节点拉选票
		for i := 0; i < raftConut; i++ {
			if i != rf.me {
				//拉选票
				go func(){
					if leader.LeaderId < 0 {
						//设置投票
						rf.elecCh <- true
					}
				}()
			}
		}
		//设置投票数量
		vote = 1
		//遍历节点拉选票
		for i := 0; i < raftConut; i++ {
			//计算投票的数量
			select {
			case ok := <- rf.elecCh:
				if ok {
					vote++
					//若选票个数，大于节点个数 / 2，则成功
					success = vote > raftConut / 2
					if success && !triggerHeartbeat{
						//开始触发心跳信号检测
						triggerHeartbeat = true
						rf.mu.Lock()
						rf.becomeLeader()
						rf.mu.Unlock()
						//由leader向其他节点发送心跳信号
						rf.heartBeat <- true
						fmt.Println(rf.me, "号节点成为了leader")
						fmt.Println("leader开始发送心跳信号")
					}
				}
			}
		}
		//做最后的检验工作
		if timeout < millisecond() || (vote > raftConut / 2 || rf.currentLeader > -1) {
			break
		}else{
			//等待操作
			select {
			case <- time.After(time.Duration(10) * time.Millisecond):
			}
		}
	}
	return success
}

/**
修改状态
 */
func (rf *Raft) becomeCandidate(){
	rf.state = 1
	rf.setTerm(rf.currentTerm + 1)
	rf.votedFor = rf.me
	rf.currentLeader = -1
}

/**
 成为领导者
 */
func (rf *Raft) becomeLeader(){
	rf.state = 2
	rf.currentLeader = rf.me
}

/**
 * leader节点发送心跳信号
 * 顺便完成数据同步
 * 看小弟有没有down
 */
func (rf *Raft) sendLeaderBeat(){
	for{
		select {
		case <- rf.heartBeat:
			rf.sendAppendEntriesImpl()
		}
	}
}

/**
 * 用于返回给leader的确认信号
 */
func (rf *Raft)sendAppendEntriesImpl() {
	//是主节点就不跑了
	if rf.currentLeader == rf.me {
		//此时是leader
		//记录确认信号的节点个数
		var success_count = 0
		//设置确认信号
		for i := 0; i < raftConut; i++ {
			if i != rf.me {
				go func(){
					//rf.heartbeatRe <- true
					//这里实际相当于客户端
					rp, err := rpc.Dial("tcp", "127.0.0.1:8080")
					if err != nil {
						log.Fatal(err)
					}
					//接收服务器返回的信息
					//接收服务端返回信息的变量
					ok := false
					err = rp.Call("Raft.Communication", Param{"Hello"}, ok)
					if err != nil {
						log.Fatal(err)
					}
					if ok {
						rf.heartbeatRe <- true
					}
				}()
			}
		}
		//再计算返回确认信号的个数
		for i := 0; i < raftConut; i++ {
			select {
			case ok := <- rf.heartbeatRe:
				if ok {
					success_count++
					if success_count > raftConut / 2 {
						fmt.Println("投票选举成功")
						log.Fatal("程序结束")
					}
				}
			}
		}
	}
}

type Param struct {
	Msg string
}

func (rf *Raft)Communication(p Param, a *bool) error{
	fmt.Println(p.Msg)
	*a = true
	return nil
}