package logger_test

import (
	"fmt"
	"sync"
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/skyterra/st-logger/logger"
)

var _ = Describe("Logger", func() {
	Context("each log level output", func() {
		It("should be succeed", func() {
			for level := logger.DEBUG; level <= logger.ERROR; level++ {
				fmt.Printf("--log level:%d\n", level)
				logger.SetLogLevel(level)
				logger.Debug("This is debug", logger.DEBUG)
				logger.Debugf("this is debug, level:%d", logger.DEBUG)

				logger.Info("this is info", logger.INFO)
				logger.Infof("this is info, level:%d", logger.INFO)

				logger.Warn("this is warn", logger.WARN)
				logger.Warnf("this is warn, level:%d", logger.WARN)

				//logger.Error("this is error", logger.ERROR)
				//logger.Errorf("this is error, level:%d", logger.ERROR)

				logger.Infof("/api/ws_main/login. timeSpent:%d", 10)
			}
		})
	})

	Context("Header field", func() {
		It("should be succeed", func() {
			logger.ShowHeader(logger.HeaderLevel, true)
			logger.ShowHeader(logger.HeaderPath, false)
			logger.Debug("hello world")
			logger.Info("hello world info")
		})
	})

	Context("concurrence", func() {
		It("should be succeed", func() {
			wg := sync.WaitGroup{}
			go func() {
				wg.Add(1)
				for i := 0; i < 100; i++ {
					logger.Infof("a hello %d", i)
				}
			}()

			go func() {
				wg.Add(1)
				for i := 0; i < 100; i++ {
					logger.Infof("b hello %d", i)
				}
			}()

			wg.Wait()
			time.Sleep(10 * time.Second)
		})
	})
})
