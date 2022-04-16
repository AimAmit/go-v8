// const _ = require('lodash');
const axios = require("axios");

const main = async (o) => {

    // console.log(_.get({"name": "amit"}, "amit"))


    const x = await axios.get("http://google.com")
    console.log(x.statusText)
    return 0;
}

// export default main;
module.exports = main
// main()