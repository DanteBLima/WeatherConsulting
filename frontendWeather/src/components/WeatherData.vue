<template>
  <div class="weather-data-card">
    <h2>Weather Data</h2>
    <div class="weather-data-content">
      <template v-if="provider === 'Open-Meteo' || provider === 'WeatherAPI'">
        <div v-for="(value, key) in weatherData" :key="key" class="weather-data-item">
          <template v-if="key === 'daily'">
            <div v-for="(dailyValue, index) in value.time" :key="index" class="daily-data">
              <h3>{{ dailyValue }}</h3>
              <div><strong>Weather Code:</strong> {{ value.weather_code[index] }}</div>
              <div><strong>Max Temp:</strong> {{ value.temperature_2m_max[index] }}°C</div>
              <div><strong>Min Temp:</strong> {{ value.temperature_2m_min[index] }}°C</div>
              <div><strong>Max Apparent Temp:</strong> {{ value.apparent_temperature_max[index] }}°C</div>
              <div><strong>Min Apparent Temp:</strong> {{ value.apparent_temperature_min[index] }}°C</div>
            </div>
          </template>
          <template v-else>
            <strong>{{ key }}:</strong> <span>{{ value }}</span>
          </template>
        </div>
      </template>
      <template v-else-if="provider === 'VisualCrossing'">
        <div v-for="(day, index) in weatherData.days" :key="index" class="daily-data">
          <h3>Day {{ index + 1 }}</h3>
          <div><strong>DateTime:</strong> {{ day.datetime }}</div>
          <div><strong>Temp Max:</strong> {{ day.tempmax }}°C</div>
          <div><strong>Temp Min:</strong> {{ day.tempmin }}°C</div>
          <div><strong>Feels Like Max:</strong> {{ day.feelslikemax }}°C</div>
          <div><strong>Feels Like Min:</strong> {{ day.feelslikemin }}°C</div>
          <div><strong>Conditions:</strong> {{ day.conditions }}</div>
          <div v-for="(hour, hourIndex) in day.hours" :key="hourIndex" class="hourly-data">
            <h4>Hour {{ hourIndex + 1 }}</h4>
            <div><strong>DateTime:</strong> {{ hour.datetime }}</div>
            <div><strong>Temp:</strong> {{ hour.temp }}°C</div>
            <div><strong>Feels Like:</strong> {{ hour.feelslike }}°C</div>
            <div><strong>Conditions:</strong> {{ hour.conditions }}</div>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    weatherData: {
      type: Object,
      required: true,
    },
    provider: {
      type: String,
      required: true,
    },
  },
};
</script>

<style scoped>
.weather-data-card {
  background: rgba(255, 255, 255, 0.2);
  padding: 70px;
  border-radius: 10px;
  margin-top: 20px;
  width: 100%;
  max-width: 600px;
  max-height: 400px; 
  overflow-y: auto;  
}

.weather-data-card h2 {
  margin-bottom: 15px;
}

.weather-data-content {
  display: flex;
  flex-direction: column;
}

.weather-data-item {
  margin-bottom: 20px;
}

.daily-data, .hourly-data {
  margin-bottom: 20px;
}

.daily-data h3, .hourly-data h4 {
  margin-bottom: 5px;
}

.daily-data div, .hourly-data div {
  margin-bottom: 3px;
}
</style>
