package znet

import "8-7/src/zinx/ziface"

type BaseRouter struct {
}

func (b *BaseRouter) PerHandle(request ziface.IRequest) {

}

func (b *BaseRouter) Handle(request ziface.IRequest) {}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {}
