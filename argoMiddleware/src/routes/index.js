import express from 'express';
import {
  indexPage,
  messagesPage,
  argoPage,
  workFlowPage,
} from '../controllers';

const indexRouter = express.Router();

indexRouter.get('/', indexPage);
indexRouter.get('/messages', messagesPage);
indexRouter.post('/argo', argoPage);
indexRouter.get('/argo/getworkflow', workFlowPage);
export default indexRouter;
