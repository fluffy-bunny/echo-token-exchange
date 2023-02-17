/*!
* Start Bootstrap - Bare v5.0.7 (https://startbootstrap.com/template/bare)
* Copyright 2013-2021 Start Bootstrap
* Licensed under MIT (https://github.com/StartBootstrap/startbootstrap-bare/blob/master/LICENSE)
*/
// This file is intentionally blank
// Use this file to add JavaScript to your project

 
function getCookieValue(name) {
    const nameString = name + "="
  
    const value = document.cookie.split(";").filter(item => {
      return item.includes(nameString)
    })
  
    if (value.length) {
      return value[0].substring(nameString.length, value[0].length)
    } else {
      return ""
    }
  }

    