import * as fs from 'fs';

/**
 * A Lambda function that returns a static string
 */
export const handler = async (event) => {
  console.log('Received event:', JSON.stringify(event, null, 2));
  // Read jokes data
  const jokesData = fs.readFileSync('./data/chuck-jokes.json', 'utf8');
  const jokes = JSON.parse(jokesData);
  // Get random joke
  const randomJoke = jokes[Math.floor(Math.random() * jokes.length)];
  // Return joke
  return {
    isBase64Encoded: false,
    statusCode: 200,
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(randomJoke),
  };
}
