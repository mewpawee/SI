import { submitWorkflow, getWorkflow } from '../utils/argoQuery';
import { verifyGoogleToken } from '../utils/googleValidate';
// const targetDomain = ['cie.kmitl.ac.th', 'kmitl.ac.th'];

export const argoPage = async (req, res) => {
  try {
    let { targetDomain } = req.query;
    targetDomain = JSON.stringify(targetDomain);
    const { token } = req.headers;
    const userId = await verifyGoogleToken(token);
    if (userId) {
      const response = await submitWorkflow(userId, targetDomain);
      res.status(200).json(response.metadata);
    } else {
      throw userId;
    }
  } catch (err) {
    res.status(200).json(err);
  }
};

export const workFlowPage = async (req, res) => {
  try {
    const { token } = req.headers;
    const { workflowId } = req.query;
    const userId = await verifyGoogleToken(token);
    if (userId) {
      const response = await getWorkflow(workflowId);
      res.status(200).json(response.status);
    } else {
      throw userId;
    }
  } catch (err) {
    res.status(200).json(err);
  }
};
