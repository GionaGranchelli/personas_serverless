<script setup lang="ts">
</script>

<template>
  <main>
    <h2>Welcome to the Persona Management System</h2>
    <p>
      The Persona Management System is a powerful tool designed to help you manage information about people efficiently.
      Whether you need to create new entries, update existing ones, or simply browse through the list of personas, our system
      provides an intuitive and user-friendly interface to handle all your needs.
    </p>
    <div>&nbsp;</div>
    <p>
      Key Features:
    </p>
    <div>&nbsp;</div>
    <ul>
      <li>Create new persona entries with ease</li>
<!--      <li>Update existing persona information</li>-->
<!--      <li>Delete personas when no longer needed</li>-->
      <li>Search for personas by ID</li>
      <li>View a comprehensive list of all personas</li>
    </ul>
    <div>&nbsp;</div>
    <p>
      Our system is built using modern serverless architecture, ensuring scalability and reliability. It leverages AWS services
      such as Lambda, API Gateway, and DynamoDB to deliver a seamless experience.
    </p>
  </main>
</template>
<style>
main {
  min-height: 100vh;
  padding: 40px;
  margin-top: 3rem;
  box-sizing: border-box;
  background: rgba(0, 0, 0, .5);
  box-shadow: 0 15px 25px rgba(0, 0, 0, .6);
  border-radius: 10px;
  max-width: 800px;
}
h2 {
  margin: 0 0 30px;
  padding-top: 40px;
  color: #fff;
  text-align: center;
}
</style>