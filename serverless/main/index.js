'use strict';

const Alexa = require('alexa-sdk');
const request = require('request');
const handlers = {
  SafeIntent: function() {
    request.get(process.env.PI_ENDPOINT, (err, res) => {
      request.post(process.env.TEXT_ENDPOINT, (err, res) => {
        this.emit(':tell', '');
      });
    });
  },
  LockHouseIntent: function() {
    request.get(process.env.MOTION_DETECTION_ON_ENDPOINT, (err, res) => {
      this.emit(':tell', 'Ok, I turned on motion detection.');
    });
  },
  UnlockHouseIntent: function() {
    request.get(process.env.MOTION_DETECTION_OFF_ENDPOINT, (err, res) => {
      this.emit(':tell', 'Welcome back, I have turned off motion detection.');
    });
  },
};

exports.handler = (event, context) => {
  const alexa = Alexa.handler(event, context);
  alexa.appId = process.env.APP_ID;
  alexa.registerHandlers(handlers);
  alexa.execute();
};
