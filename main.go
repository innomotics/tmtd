/*
Copyright © 2024 Harald Müller <harald.mueller@evosoft.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/wot-oss/tmtd/cmd"
	"github.com/wot-oss/tmtd/internal"
	"github.com/wot-oss/tmtd/internal/config"
)

func init() {
	config.InitConfig()
	config.InitViper()
	internal.InitLogging()
}

func main() {
	// opts := &slog.HandlerOptions{
	// 	Level: slog.LevelDebug,
	// }
	// handler := slog.NewJSONHandler(os.Stdout, opts)
	// logger := slog.New(handler)
	// slog.SetDefault(logger)
	cmd.Execute()
}
