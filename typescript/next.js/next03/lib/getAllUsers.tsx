import React from "react";

export const getAllUsers = async () => {
  const res = await fetch("http://jsonplaceholder.typicode.com/users");
  if (!res.ok) throw new Error("Failed to fetch data");
  return res.json();
};
