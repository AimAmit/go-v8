// const _ = require('lodash');
// const axios = require("axios");
const chalk = require('chalk');


const main = async (o) => {

    // console.log(_.get({"name": "amit"}, "amit"))


    // const x = await axios.get("http://google.com")
    console.log("x.statusText")
    console.log(chalk.blue('Hello world!'));

    return 10;
}

// export default main;
// module.exports = main
main()