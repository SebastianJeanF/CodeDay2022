const base_url =  "http://api.openweathermap.org/data/2.5/weather?";
const city_name = input("City name?: ");
const API_KEY = "64b38d3f58717f4e8c8c6a8033addac1";

const final_url_geocoding = base_url + "appid=" + API_KEY + "&q=" + city_name;

fetch(final_url_geocoding, {
  "method": "GET"

})
.then(response => {
  console.log(response);
})
.catch(err => {
  console.log(err);
});

console.log()