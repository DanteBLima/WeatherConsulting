<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import WeatherData from './components/WeatherData.vue';

const providers = ['Open-Meteo', 'WeatherAPI', 'VisualCrossing'];
const options = ['Current Weather', '7 days forecast', 'Start to End date'];
const selectedProvider = ref('');
const selectedOption = ref('');
const cityName = ref('');
const weatherData = ref(null);
const startDate = ref('');
const endDate = ref('');
const formVisible = ref(true);

const numPredictions = ref(1);
const predictions = ref([
  { provider: '', option: '', city: '', startDate: '', endDate: '' }
]);

const addPrediction = () => {
  if (numPredictions.value > predictions.value.length) {
    predictions.value.push({ provider: '', option: '', city: '', startDate: '', endDate: '' });
  } else {
    predictions.value = predictions.value.slice(0, numPredictions.value);
  }
};

const provideInfo = async () => {
  try {
    for (const prediction of predictions.value) {
      let start = '';
      let end = '';

      if (prediction.option === 'Start to End date') {
        start = prediction.startDate;
        end = prediction.endDate;
      }

      let typeParam = '';
      if (prediction.provider === 'Open-Meteo') {
        typeParam = prediction.option === 'Current Weather' ? 'current' :
                    prediction.option === '7 days forecast' ? 'dynamic' : 
                    'historical';
      } else if (prediction.provider === 'WeatherAPI') {
        typeParam = prediction.option === 'Current Weather' ? 'current3' :
                    prediction.option === '7 days forecast' ? 'dynamic3' : 
                    'historical3';
      } else if (prediction.provider === 'VisualCrossing') {
        typeParam = prediction.option === 'Current Weather' ? 'current2' :
                    prediction.option === '7 days forecast' ? 'dynamic2' : 
                    'historical2';
      }

      let queryParams = '';
      if (prediction.provider === 'Open-Meteo') {
        const geoResponse = await axios.get('https://api.weatherapi.com/v1/current.json', {
          params: {
            key: 'b26291ee7a8c49ed8e114442242606',
            q: prediction.city,
          },
        });

        const { lat, lon } = geoResponse.data.location;

        queryParams = prediction.option === 'Start to End date' 
          ? `?type=${typeParam}&lat=${lat}&lon=${lon}&start=${start}&end=${end}` 
          : `?type=${typeParam}&lat=${lat}&lon=${lon}`;
      } else {
        queryParams = prediction.option === 'Start to End date' 
          ? `?type=${typeParam}&city=${prediction.city}&start=${start}&end=${end}` 
          : `?type=${typeParam}&city=${prediction.city}`;
      }

      const baseUrl = 'http://192.168.15.39:8080/weather';
      const weatherResponse = await axios.get(baseUrl + queryParams);

      prediction.weatherData = weatherResponse.data;
    }
    formVisible.value = false;
  } catch (error) {
    console.error('Erro ao obter dados meteorológicos: ', error);
  }
};

const ModifyOptionChange = (predictionIndex) => {
  predictions.value[predictionIndex].startDate = '';
  predictions.value[predictionIndex].endDate = '';
};

onMounted(() => {
  
});
</script>

<template>
  <div class="background">
    <img src = "/public/weatherTitle.png" alt="logo" class="logo">
    <div class="container" v-if="formVisible">
      <div class="dropdown-container">
        <select v-model="numPredictions" class="dropdown" @change="addPrediction">
          <option value="1">1 Previsão</option>
          <option value="2">2 Previsões</option>
          <option value="3">3 Previsões</option>
        </select>
      </div>
      <div class="predictions-container">
        <div v-for="(prediction, index) in predictions" :key="index" class="prediction-box">
          <div class="dropdown-container">
            <select v-model="prediction.provider" class="dropdown">
              <option disabled value="">Select Provider</option>
              <option v-for="provider in providers" :key="provider" :value="provider">{{ provider }}</option>
            </select>
            <select v-model="prediction.option" class="dropdown" @change="() => ModifyOptionChange(index)">
              <option disabled value="">Select Option</option>
              <option v-for="option in options" :key="option" :value="option">{{ option }}</option>
            </select>
          </div>
          <input type="text" v-model="prediction.city" class="input-field" placeholder="City Name" />
          <div v-if="prediction.option === 'Start to End date'" class="date-inputs">
            <input type="text" v-model="prediction.startDate" class="input-field" placeholder="YYYY-MM-DD (Start Date)" />
            <input type="text" v-model="prediction.endDate" class="input-field" placeholder="YYYY-MM-DD (End Date)" />
          </div>
        </div>
      </div>
      <button class="get-started-button" @click="provideInfo">PROVIDE</button>
    </div>
    
    <div class="results-container" v-if="!formVisible">
      <div v-for="(prediction, index) in predictions" :key="index">
        <WeatherData v-if="prediction.weatherData" :weatherData="prediction.weatherData" :provider="prediction.provider" />
      </div>
    </div>
  </div>
</template>

<style>

.logo {
  position: absolute;
  top: 0%;
  left: 50%;
  transform: translateX(-50%);
  z-index: 1;
}


body {
  
  margin: 0;
  font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
  overflow: hidden;
  height: 100vh;
  background-size: cover;
  background-position: center center;
  background-image: url("img.jpg");
}

.background {
  
  width: 100%;
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-size: cover;
  background-position: center;
  top: 50%;
}

.container {
  background: rgba(255, 255, 255, 0.1); 
  border: 2px solid rgba(255, 255, 255, 0.2);
  position: absolute;
  z-index: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  border-radius: 10px;
}

.dropdown-container {
  
  display: flex;
  justify-content: space-between;
  width: 100%;
  margin-bottom: 20px;
}

.dropdown {
  
  background: transparent;
  flex: 1;
  margin: 0 10px;
  padding: 10px;
  font-size: 16px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.input-field {
  
  width: calc(100% - 20px);
  padding: 10px;
  font-size: 16px;
  margin-bottom: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.date-inputs {
  display: flex;
  justify-content: space-between;
  width: 100%;
}

.date-inputs .input-field {
  flex: 1;
  margin: 0 5px;
}

.get-started-button {
  font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
  padding: 10px 20px;
  font-size: 18px;
  background-color: #1a0e55;
  color: #fff;
  border: none;
  border-radius: 20px;
  cursor: pointer;
}

.get-started-button:hover {
  background-color: #120844;
}

.predictions-container {
  
  display: flex;
  justify-content: space-between;
  width: 100%;
  flex-wrap: wrap;
}

.prediction-box {
  flex: 1 1 45%;
  background: rgba(255, 255, 255, 0.1); 
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 10px;
  padding: 15px;
  margin: 10px;
}

.results-container {
  top:30%;
  position:absolute;
  display: flex;
  justify-content: center;
  width: 100%;
  flex-wrap: wrap;
  margin: 0 auto;
  
}

.results-container > div {
  
  justify-content: center;
  margin: 0 auto;
  flex: center;
  margin: 10px;
}
</style>
