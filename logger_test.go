package logger_test

import (
	"context"
	"sync"

	. "github.com/onsi/ginkgo"
	"github.com/skyterra/logger"
)

var _ = Describe("Logger", func() {
	Context("set log level", func() {
		FIt("should be succeed", func() {
			logger.SetLevel("debug")
			logger.SetProjectName("ST-Logger")
			logger.SetSrcFolder("logger")
			logger.Debug(context.TODO(), "this is debug")

			ctx := context.WithValue(context.TODO(), logger.RequestID, "aaa")
			logger.Info(ctx, "this is info")

			logger.SetLevel("info")
			logger.Debug(context.TODO(), "this is debug")

			logger.Error(context.TODO(), "this is test error")
		})
	})

	Context("set project", func() {
		It("should be succeed", func() {
			logger.SetProjectName("ST-Logger")
			logger.Debug(context.TODO(), "this is debug")
		})
	})

	Context("set src folder", func() {
		It("should be succeed", func() {
			logger.SetSrcFolder("st-logger")
			logger.Debug(context.TODO(), "this is test")
		})
	})

	Context("concurrence", func() {
		It("should be succeed", func() {
			wg := sync.WaitGroup{}
			wg.Add(2)
			go func() {
				for i := 0; i < 200; i++ {
					logger.Infof(context.TODO(), "a hello %d", i)
				}
				wg.Done()
			}()

			go func() {
				for i := 0; i < 200; i++ {
					logger.Infof(context.TODO(), "b hello %d", i)
				}
				wg.Done()
			}()

			wg.Wait()
		})
	})
})
