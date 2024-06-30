<template>
  <div>
    <h1>Create Person</h1>
    <form @submit.prevent="createPerson">
      <div>
        <label for="firstName">First Name:</label>
        <input type="text" v-model="person.firstName" required />
      </div>
      <div>
        <label for="lastName">Last Name:</label>
        <input type="text" v-model="person.lastName" required />
      </div>
      <div>
        <label for="phoneNumber">Phone Number:</label>
        <input type="text" v-model="person.phoneNumber" required />
      </div>
      <div>
        <label for="address">Address:</label>
        <input type="text" v-model="person.address" required />
      </div>
      <button type="submit">Create</button>
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from 'vue';
import axios from 'axios';
import config from '../config';

interface Person {
  firstName: string;
  lastName: string;
  phoneNumber: string;
  address: string;
}

export default defineComponent({
  name: 'CreatePersona',
  setup() {
    const person = reactive<Person>({
      firstName: '',
      lastName: '',
      phoneNumber: '',
      address: '',
    });

    const createPerson = async () => {
      try {
        const response = await axios.post(`${config.apiEndpoint}/personas`, person);
        console.log(response.data);
        alert('Person created successfully!');
        Object.assign(person, { firstName: '', lastName: '', phoneNumber: '', address: '' });
      } catch (error) {
        console.error(error);
        alert('Failed to create person.');
      }
    };

    return {
      person,
      createPerson,
    };
  },
});
</script>
