import { OAuth2Client } from 'google-auth-library';

const CLIENT_ID = '647735027802-876e7ib9pi85sbtcq2sa626c1rigcnqi.apps.googleusercontent.com';
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
