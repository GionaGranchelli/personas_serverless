<template>
  <div class="search-wrapper">
    <h2>Search Person by ID</h2>
    <form @submit.prevent="searchPerson">
      <div class="control">
        <input type="text" v-model="personId" required />
        <label for="id">ID:</label>
      </div>
      <div>
        &nbsp;
      </div>
      <button type="submit">Search</button>
    </form>
    <div v-if="person">
      <h2>Person Details</h2>
      <p><strong>First Name:</strong> {{ person.first_name }}</p>
      <p><strong>Last Name:</strong> {{ person.last_name }}</p>
      <p><strong>Phone Number:</strong> {{ person.phone_number }}</p>
      <p><strong>Address:</strong> {{ person.address }}</p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, inject, ref } from 'vue'
import axios from 'axios';

interface Person {
  first_name: string;
  last_name: string;
  phone_number: string;
  address: string;
}

export default defineComponent({
  name: 'SearchPersona',
  setup() {
    const apiEndpoint = inject<string>('apiEndpoint');
    const personId = ref('');
    const person = ref<Person | null>(null);

    const searchPerson = async () => {
      try {
        const response = await axios.get(`${apiEndpoint}/personas/${personId.value}`);
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
<style>
.search-wrapper {
  padding: 40px;
  margin-top: 3rem;
  min-height: 100vh;
  box-sizing: border-box;
  background: rgba(0,0,0,.5);
  box-shadow: 0 15px 25px rgba(0,0,0,.6);
  border-radius: 10px;

}
.search-wrapper h2 {
  margin: 0 0 30px;
  padding-top: 40px;
  color: #fff;
  text-align: center;
}
.control {
  position: relative;
}

.search-wrapper .control input {
  width: 100%;
  padding: 10px 0;
  font-size: 16px;
  color: #fff;
  margin-bottom: 30px;
  border: none;
  border-bottom: 1px solid #fff;
  outline: none;
  background: transparent;
}
.search-wrapper .control label {
  position: absolute;
  top:0;
  left: 0;
  padding: 10px 0;
  font-size: 16px;
  color: #fff;
  pointer-events: none;
  transition: .5s;
}

.search-wrapper .control input:focus ~ label,
.search-wrapper .control input:valid ~ label {
  top: -20px;
  left: 0;
  color: hsla(160, 100%, 37%, 1);
  font-size: 12px;
}

.search-wrapper form a {
  position: relative;
  display: inline-block;
  padding: 10px 20px;
  color: hsla(160, 100%, 37%, 1);
  font-size: 16px;
  text-decoration: none;
  text-transform: uppercase;
  overflow: hidden;
  transition: .5s;
  margin-top: 40px;
  letter-spacing: 4px
}

</style>