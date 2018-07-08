function log(text) {
    return new Promise(resolve => {
        setTimeout(() => {
            console.log(text);
            resolve();
        }, 250);
    })
}


function logTimes(count) {
    let sequence = Promise.resolve();

    for (let i = 0; i < count; i += 1) {
        sequence = sequence.then(() => log(i));
    }
    return sequence;
}


logTimes(5)
    .then(
        () => console.log('Done'),
        error => console.error(error)
    );