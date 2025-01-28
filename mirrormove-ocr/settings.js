"use strict";

const ENV = process.env;

exports.JWT = {
	secretKey: ENV["SECRETKEY"]
};