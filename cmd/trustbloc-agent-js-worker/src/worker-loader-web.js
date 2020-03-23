/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

export function loadWorker(pending, notifications, paths) {
    const wasmJS = paths.wasmJS
    const wasm = paths.wasm
    const workerJS = paths.dir + "/worker-impl-web.js"
    const worker = new Worker(workerJS + "?wasmJS=" + wasmJS + "&wasm=" + wasm)
    worker.onmessage = e => {
        const result = e.data
        if (result.topic){
            if (notifications.get(result.topic)) {
                notifications.get(result.topic)(result)
            }  else if (notifications.get("all")){
                notifications.get("all")(result)
            } else {
                console.log("no subscribers found for this topic", result.topic)
            }
          return
        }
        const cb = pending.get(result.id)
        pending.delete(result.id)
        cb(result)
    }
    worker.onerror = e => {
        throw new Error("trustblocagent: failed to load worker: " + e.message)
    }
    return worker
}
