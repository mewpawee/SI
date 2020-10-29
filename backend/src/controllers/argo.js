import { submitWorkflow, getWorkflow } from '../utils/argoQuery';

// const targetDomain = ['cie.kmitl.ac.th', 'kmitl.ac.th'];

export const argoPage = async (req, res) => {
  try {
    const { userId } = req.query;
    let { targetDomain } = req.query;
    targetDomain = JSON.stringify(targetDomain);
    // console.log(userId);
    const response = await submitWorkflow(userId, targetDomain);
    res.status(200).json(response.metadata);
  } catch (err) {
    res.status(200).json(err);
  }
};

export const workFlowPage = async (req, res) => {
  try {
    const { workflowId } = req.query;
    const response = await getWorkflow(workflowId);
    res.status(200).json(response.status);
  } catch (err) {
    res.status(200).json(err);
  }
};
