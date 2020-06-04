import { Response } from "../models/quote";

const url = "http://localhost:3333";

export async function getQuote(): Promise<Response> {
  const response = await fetch(`${url}/qotd`);
  return await response.json();
}
