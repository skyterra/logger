package logger_test

import (
	"sync"

	. "github.com/onsi/ginkgo"
	"github.com/skyterra/st-logger/logger"
)

var _ = Describe("Logger", func() {
	Context("set log level", func() {
		It("should be succeed", func() {
			logger.SetLevel("debug")
			logger.SetProjectName("ST-Logger")
			logger.SetSrcFolder("st-logger")
			logger.Debug("this is debug")
			logger.Info("this is info")

			logger.SetLevel("info")
			logger.Debug("this is debug")

			logger.Error("this is test error")
		})
	})

	Context("set project", func() {
		It("should be succeed", func() {
			logger.SetProjectName("ST-Logger")
			logger.Debug("this is debug")
		})
	})

	Context("set src folder", func() {
		It("should be succeed", func() {
			logger.SetSrcFolder("st-logger")
			logger.Debug("this is test")
		})
	})

	Context("concurrence", func() {
		It("should be succeed", func() {
			wg := sync.WaitGroup{}
			wg.Add(2)
			go func() {
				for i := 0; i < 200; i++ {
					logger.Infof("a hello %d", i)
				}
				wg.Done()
			}()

			go func() {
				for i := 0; i < 200; i++ {
					logger.Infof("b hello %d", i)
				}
				wg.Done()
			}()

			wg.Wait()
		})
	})
})
