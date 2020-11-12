import { submitWorkflow, getWorkflow } from '../utils/argoQuery';
import { verifyGoogleToken } from '../utils/validateToken';
// const targetDomain = ['cie.kmitl.ac.th', 'kmitl.ac.th'];

export const argoPage = async (req, res) => {
  try {
    // console.log(req.body);
    const targetEndpoints = JSON.stringify(req.body);
    let { token } = req.headers;
    if (token) {
      if (token.startsWith('Bearer ')) {
        token = token.slice(7, token.length).trimLeft();
      } else {
        throw new Error("can't find Bearer");
      }
    }
    const sub = await verifyGoogleToken(token);
    if (sub) {
      const response = await submitWorkflow(sub, targetEndpoints);
      return res.status(200).json(response.metadata);
    }
  } catch (err) {
    return res
      .status(401)
      .json({ error: err.name, error_description: err.message });
  }
};

export const workFlowPage = async (req, res) => {
  try {
    const { token } = req.headers;
    if (token) {
      if (token.startsWith('Bearer ')) {
        token = token.slice(7, token.length).trimLeft();
      } else {
        throw new Error("can't find Bearer");
      }
    }
    const { workflowId } = req.query;
    const sub = await verifyGoogleToken(token);
    if (sub) {
      const response = await getWorkflow(workflowId);
      return res.status(200).json(response.status);
    } else {
      throw new Error('cant get workflow');
    }
  } catch (err) {
    return res
      .status(401)
      .json({ error: err.name, error_description: err.message });
  }
};
