<template>
  <div>
    <h1>Search Person by ID</h1>
    <form @submit.prevent="searchPerson">
      <div>
        <label for="id">ID:</label>
        <input type="text" v-model="personId" required />
      </div>
      <button type="submit">Search</button>
    </form>
    <div v-if="person">
      <h2>Person Details</h2>
      <p><strong>First Name:</strong> {{ person.firstName }}</p>
      <p><strong>Last Name:</strong> {{ person.lastName }}</p>
      <p><strong>Phone Number:</strong> {{ person.phoneNumber }}</p>
      <p><strong>Address:</strong> {{ person.address }}</p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import axios from 'axios';
import config from '../config';

interface Person {
  firstName: string;
  lastName: string;
  phoneNumber: string;
  address: string;
}

export default defineComponent({
  name: 'SearchPersona',
  setup() {
    const personId = ref('');
    const person = ref<Person | null>(null);

    const searchPerson = async () => {
      try {
        const response = await axios.get(`${config.apiEndpoint}/personas/${personId.value}`);
        person.value = response.data;
      } catch (error) {
        console.error(error);
        alert('Failed to retrieve person.');
      }
    };

    return {
      personId,
      person,
      searchPerson,
    };
  },
});
</script>
