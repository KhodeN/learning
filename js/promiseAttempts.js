let counter = 0;
const item = {id: 1, name: 'item'};

function getItem() {
    counter += 1;
    console.log('try to get item', counter);

    return new Promise((resolve, reject) => {
        setTimeout(() => {
            if (counter > 4) {
                resolve(item)
            } else {
                reject('not yet')
            }
        }, 250);
    });
}


function tryToGetItem(tryCountLimit) {
    if (tryCountLimit === 0) {
        return Promise.reject('attempts exceed');
    }
    return getItem()
        .catch(() => tryToGetItem(tryCountLimit - 1));
}


tryToGetItem(5)
    .then(
        item => console.log(item),
        error => console.error(error),
    );