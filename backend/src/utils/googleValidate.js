import { OAuth2Client } from 'google-auth-library';

const CLIENT_ID = '123';
const client = new OAuth2Client(CLIENT_ID);

export const verifyGoogleToken = async (token) => {
  const ticket = await client.verifyIdToken({
    idToken: token,
    audience: CLIENT_ID,
  });
  const payload = ticket.getPayload();
  const userid = payload.sub;
  return userid;
};
