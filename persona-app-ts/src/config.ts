interface Config {
  apiEndpoint: string;
}

let config: Config = {
  apiEndpoint: '',
};

async function fetchConfig() {
  try {
    const response = await fetch('/api-config.json');
    const data = await response.json();
    config.apiEndpoint = data.apiUrl;
  } catch (error) {
    console.error('Error fetching API config:', error);
  }
}

export default { ...config, fetchConfig };
