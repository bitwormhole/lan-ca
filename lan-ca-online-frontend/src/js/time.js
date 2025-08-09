// time.js



function formatTime(time_in_ms) {
    let date = new Date();
    date.setTime(time_in_ms)
    return date.toLocaleString()
}




export const Time = {
    formatTime
}
