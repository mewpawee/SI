import axios from 'axios';

const url = 'https://www.googleapis.com/oauth2/v3/userinfo';

export const verifyGoogleToken = async (token) => {
  const params = {
    access_token: token,
  };
  try {
    const response = await axios.get(url, { params });
    const { data } = response;
    return data.sub;
  } catch (err) {
    throw new Error("Unable to verify a google's access token.");
  }
};
