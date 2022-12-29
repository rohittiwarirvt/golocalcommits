// signle async await
const sleep = ms => new Promise(r => setTimeout(r, ms));

const longRunningTask = async () => {
    // Simulate a workload.
    sleep(3000)
    return Math.floor(Math.random() * Math.floor(100))
}

// simple await
let r = await longRunningTask()
console.log(r)








// Promise.all
const [a, b, c] = await Promise.all(longRunningTask(), longRunningTask(), longRunningTask())
console.log(a, b, c)


// Promise.race()

const one = async () => {
    // Simulate a workload.
    sleep(Math.floor(Math.random() * Math.floor(2000)))
    return 1
}

const two = async () => {
    // Simulate a workload.
    sleep(Math.floor(Math.random() * Math.floor(1000)))
    sleep(Math.floor(Math.random() * Math.floor(1000)))
    return 2
}

r = await Promise.race(one(), two())
console.log(r)