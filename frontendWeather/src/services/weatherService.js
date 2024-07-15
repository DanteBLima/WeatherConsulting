import api from '../api'

export const getCurrentWeather = async (lat, lon) => {

  try {
    const response = await api.get('/weather', {
      params: {
        type: 'current',
        lat: lat,
        lon: lon,
      },
    });
    return response.data;
  } catch (error) {
    console.error('Error trying to get current weather: ', error);
    throw errpr;
  }
};


export const getDynamicWeather = async (lat, lon) => {
  try {
    const response = await api.get('/weather', {
      params: {
        type: 'dynamic',
        lat: lat,
        lon: lon,
      },
    });
    return response.data;
  } catch (error) {
    console.error('Error trying to get dynamic weather: ', error);
    throw error;
  }
};


export const getHistoricalWeather = async (lat, lon, startDate, endDate) => {
  try {
    const response = await api.get('/weather', {
      params: {
        type: 'historical',
        lat: lat,
        lon: lon,
        startDate: startDate,
        endDate: endDate,
      },
    });
    return response.data;
  } catch (error) {
    console.error('Error trying to get historical weather: ', error);
    throw error;
  }
};




