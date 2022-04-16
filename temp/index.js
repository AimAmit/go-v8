// f = (a: number[], b: number, c: string) => number
const math = require('./math')

const main = (o) => {
    console.log(math.sum(10, 20))
    console.log(JSON.stringify(o))
    // console.log("js: ", Object.entries(o))
    // console.log(Object.entries(o.array))
    //a.reduce((c, acc) => acc + c, b) + c.length;
    return 5;
};
// let result = main([10, 20], 23, 345);
// console.log(result)
// result;

// let d = performance.now()
// main([10, 20], 23, '345')
// console.log((performance.now() - d)*1000)

module.exports = main
// main()