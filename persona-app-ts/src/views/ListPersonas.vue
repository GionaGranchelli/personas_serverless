<template>
  <div>
    <h1>List of Persons</h1>
    <ul>
      <li v-for="person in persons" :key="person.id">{{ person.firstName }} {{ person.lastName }} - {{ person.phoneNumber }} - {{ person.address }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import axios from 'axios';
import config from '../config';

interface Person {
  id: string;
  firstName: string;
  lastName: string;
  phoneNumber: string;
  address: string;
}

export default defineComponent({
  name: 'ListPersonas',
  setup() {
    const persons = ref<Person[]>([]);
    const fetchPersons = async () => {
      try {
        const response = await axios.get(`${config.apiEndpoint}/personas`);
        persons.value = response.data;
      } catch (error) {
        console.log('Failed to retrieve persons.');
        console.error(error);
        alert('Failed to retrieve persons.');
        persons.value = [];
      }
    };

    onMounted(fetchPersons);

    return {
      persons,
    };
  },
});
</script>
