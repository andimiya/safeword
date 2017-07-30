'use strict';

const Alexa = require('alexa-sdk');
const request = require('request');
const handlers = {
  LaunchRequest: function() {
    this.emit('SafeIntent');
  },
  SafeIntent: function() {
    request.get(process.env.VIDEO_ENDPOINT, (err, res) => {
      if (err) {
        console.log('ERROR VIDEO', err);
      }
      this.emit(':tell', 'Andrea is a total noob');
    });
    request.post(process.env.TEXT_ENDPOINT, (err, res) => {
      if (err) {
        console.log('ERROR TEXT', err);
      }
    });
  },
};

exports.handler = (event, context) => {
  const alexa = Alexa.handler(event, context);
  alexa.registerHandlers(handlers);
  alexa.execute();
};
