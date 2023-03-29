/*
 *
 * Copyright 2023 puzzlestaticserver authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Overload() == nil {
		log.Println("Loaded .env file")
	}

	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		log.Fatal("Missing SERVICE_PORT environment variable")
	}

	path := os.Getenv("BASE_PATH")
	if path == "" {
		log.Println("BASE_PATH not found, using current directory as default")
	}

	http.Handle("/", http.FileServer(http.Dir(path)))

	log.Println(http.ListenAndServe(cleanPort(port), nil))
}

func cleanPort(port string) string {
	if port[0] != ':' {
		port = ":" + port
	}
	return port
}
