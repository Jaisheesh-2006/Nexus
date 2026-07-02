import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
    stages: [
        { duration: '10s', target: 20 }, // Ramp up to 20 users over 10 seconds
        { duration: '30s', target: 50 }, // Stay at 50 users for 30 seconds
        { duration: '10s', target: 0 },  // Ramp down to 0 users
    ],
};

export default function () {
    // Your GraphQL Endpoint
    const url = 'http://localhost:8000/graphql';

    // The GraphQL query we are testing
    const payload = JSON.stringify({
        query: `
      query {
        products {
          id
          name
          price
        }
      }
    `,
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    const res = http.post(url, payload, params);

    // Check if the request was successful
    check(res, {
        'status is 200': (r) => r.status === 200,
        'no graphql errors': (r) => !r.body.includes('errors'),
    });

    sleep(1);
}
