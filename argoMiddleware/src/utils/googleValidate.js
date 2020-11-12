import { OAuth2Client } from 'google-auth-library';

const CLIENT_ID = '622455844449-11312eea0b1ogdpe8htmbr9ghracp6mn.apps.googleusercontent.com';
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
