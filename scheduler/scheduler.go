// The MIT License (MIT)
//
// Copyright (c) 2014 AT&T
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package scheduler

import (
	"github.com/att-innovate/charmander-scheduler/manager"
	"github.com/att-innovate/charmander-scheduler/mesosproto"
)

type Scheduler struct {

	// Invoked when the scheduler successfully registers with a Mesos master.
	Registered func(manager manager.Manager, frameworkId string)

	// Hook to let scheduler overwrite Task attributes (mem, cpus, nodetype, nodename ..)
	OverwriteTaskAttributes func(manager manager.Manager, taskrequest *manager.Task)

	// Invoked when resources have been offered to this framework.
	ResourceOffers func(manager.Manager, []*mesosproto.Offer)

}
