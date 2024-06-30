interface Config {
  apiEndpoint: string;
}
export type { Config };

const config: Config = {
  apiEndpoint: '',
};

async function fetchConfig() {
  try {
    const response = await fetch('/api-config.json');
    const data = await response.json();
    console.log(`here`)
    console.log(data)
    config.apiEndpoint = data.apiUrl;
    return config;
  } catch (error) {
    console.error('Error fetching API config:', error);
  }
}

export default { ...config, fetchConfig};
