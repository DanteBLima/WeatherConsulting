# WeatherConsulting
Weather forecast application made using Golang for the backend and vuejs in the frontend.


Weather forecast application made using Golang for the backend and vuejs in the frontend.

This application allows you to check the weather forecast in different ways and with different sources.

In order to test this application you need to clone the repository and get your own API key for each provider request you would like to make, 
since the api's key contained in these files are already deactivated and will not work for you.

The api keys you need are from the providers below, they are all free:
- [Open-Meteo](https://open-meteo.com)
- [Visual Crossing](https://www.visualcrossing.com)
- [WeatherAPI](https://api.weatherapi.com)

If you do not wish to use the interface made for this project, follow the instructions below to test the application using only the backend:
- Clone the backend repository
- Change the address in the router.go file to start the server on your machine or desired address
- Initialize it
- use " {address}/weather " as base url for the requests

Avaiable requests for provider 1:
- {address}/weather?type=current&lat={inform latitude}&lon={inform longitude}
- {address}/weather?type=dynamic&lat={inform latitude}&lon={inform longitude}
- {address}/weather?type=historical&lat={inform latitude}&lon={inform longitude}&start={inform start date format YYYY-MM-DD}&end={inform end date format YYYY-MM-DD}

Avaible requests for provider 2:
- {address}/weather?type=current2&city={inform city name}
- {address}/weather?type=dynamic2&city={inform city name}
- {address}/weather?type=historical2&city={inform city name}&start={inform start date format YYYY-MM-DD}&end={inform end date format YYYY-MM-DD}

Avaible request for provider 3:
- {address}/weather?type=current3&city={inform city name}
- {address}/weather?type=dynamic3&city={inform city name}


[Postman documentation](https://documenter.getpostman.com/view/34272091/2sA3JRYys5#5df5a29a-e4b0-4101-83c1-515bafa94796)
