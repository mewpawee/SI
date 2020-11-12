import { startArgo } from './argoApi';

export const getWorkflow = async (workflowId) => {
  const url = `/api/v1/workflows/argo/${workflowId}`;
  const body = {};
  return startArgo('GET', url, body);
};

export const submitWorkflow = async (userId, targetEndpoints) => {
  const url = '/api/v1/workflows/argo';
  const body = {
    namespace: 'argo',
    serverDryRun: false,
    workflow: {
      apiVersion: 'argoproj.io/v1alpha1',
      kind: 'Workflow',
      metadata: { generateName: 'scan-' },
      spec: {
        entrypoint: 'scan-service',
        imagePullSecrets: [{ name: 'registry-sec' }],
        arguments: {
          parameters: [
            { name: 'target-domain', value: `${targetEndpoints}` },
            { name: 'userId', value: `${userId}` },
          ],
        },
        serviceAccountName: 'argo',
        templates: [
          {
            name: 'scan-service',
            dag: {
              tasks: [
                {
                  name: 'scan-dnsmap',
                  template: 'dnsmap',
                  arguments: {
                    parameters: [
                      { name: 'target-domain', value: '{{item.name}}' },
                    ],
                  },
                  withParam: '{{workflow.parameters.target-domain}}',
                },
                {
                  name: 'test',
                  dependencies: ['scan-dnsmap'],
                  template: 'echo',
                  arguments: {
                    parameters: [
                      { name: 'fileLocation', value: '{{item.fileLocation}}' },
                    ],
                  },
                  withParam: '{{tasks.scan-dnsmap.outputs.parameters}}',
                },
                {
                  name: 'scan-genReport',
                  dependencies: ['test'],
                  template: 'gen-report',
                },
              ],
            },
          },
          {
            name: 'gen-report',
            inputs: {
              artifacts: [
                {
                  name: 'reportDnsmap',
                  path: '/tmp/',
                  s3: {
                    endpoint: 'fra1.digitaloceanspaces.com',
                    bucket: 'testv1',
                    key:
                      '{{workflow.parameters.userId}}/reports/{{workflow.name}}/src/dnsmap/',
                    accessKeySecret: { name: 'artifact-key', key: 'accessKey' },
                    secretKeySecret: { name: 'artifact-key', key: 'secretKey' },
                  },
                },
              ],
            },
            container: {
              image: 'registry.digitalocean.com/sec/report:v01',
              command: ['sh', '-c'],
              args: [
                'enscript /tmp/*.txt --output=- | ps2pdf -sOwnerPassword="1234" -sUserPassword="1234" - > "/tmp/{{workflow.name}}-report.pdf";',
              ],
            },
            outputs: {
              artifacts: [
                {
                  name: 'report',
                  globalName: 'report-location',
                  path: '/tmp/{{workflow.name}}-report.pdf',
                  archive: { none: {} },
                  s3: {
                    endpoint: 'fra1.digitaloceanspaces.com',
                    bucket: 'testv1',
                    key:
                      '{{workflow.parameters.userId}}/reports/{{workflow.name}}/{{workflow.name}}-report.pdf',
                    accessKeySecret: { name: 'artifact-key', key: 'accessKey' },
                    secretKeySecret: { name: 'artifact-key', key: 'secretKey' },
                  },
                },
              ],
            },
          },
          {
            name: 'echo',
            inputs: {
              parameters: [{ name: 'fileLocation' }],
              artifacts: [
                {
                  name: 'message',
                  path: '/tmp/',
                  s3: {
                    endpoint: 'fra1.digitaloceanspaces.com',
                    bucket: 'testv1',
                    key:
                      '{{workflow.parameters.userId}}/artifacts/{{workflow.name}}/raw/dnsmap/',
                    accessKeySecret: { name: 'artifact-key', key: 'accessKey' },
                    secretKeySecret: { name: 'artifact-key', key: 'secretKey' },
                  },
                },
              ],
            },
            container: {
              image: 'alpine:3.7',
              command: ['sh', '-c'],
              args: [
                'grep -Eo "([0-9]{1,3}\\.){3}[0-9]{1,3}" "/tmp/output-dnsmap-{{inputs.parameters.fileLocation}}.txt" > "/tmp/filtered-dnsmap-{{inputs.parameters.fileLocation}}.txt";',
              ],
            },
            outputs: {
              artifacts: [
                {
                  name: 'filteredDnsmap',
                  path:
                    '/tmp/filtered-dnsmap-{{inputs.parameters.fileLocation}}.txt',
                  archive: { none: {} },
                  s3: {
                    endpoint: 'fra1.digitaloceanspaces.com',
                    bucket: 'testv1',
                    key:
                      '{{workflow.parameters.userId}}/artifacts/{{workflow.name}}/filtered/dnsmap/filtered-dnsmap-{{inputs.parameters.fileLocation}}.txt',
                    accessKeySecret: { name: 'artifact-key', key: 'accessKey' },
                    secretKeySecret: { name: 'artifact-key', key: 'secretKey' },
                  },
                },
                {
                  name: 'reportDnsmap',
                  path:
                    '/tmp/output-dnsmap-{{inputs.parameters.fileLocation}}.txt',
                  archive: { none: {} },
                  s3: {
                    endpoint: 'fra1.digitaloceanspaces.com',
                    bucket: 'testv1',
                    key:
                      '{{workflow.parameters.userId}}/reports/{{workflow.name}}/src/dnsmap/output-dnsmap-{{inputs.parameters.fileLocation}}.txt',
                    accessKeySecret: { name: 'artifact-key', key: 'accessKey' },
                    secretKeySecret: { name: 'artifact-key', key: 'secretKey' },
                  },
                },
              ],
            },
          },
          {
            name: 'dnsmap',
            inputs: { parameters: [{ name: 'target-domain' }] },
            container: {
              image: 'registry.digitalocean.com/sec/dnsmap:v1',
              command: ['sh', '-c'],
              args: [
                './dnsmap "{{inputs.parameters.target-domain}}" -r "./tmp/output-dnsmap-{{inputs.parameters.target-domain}}.txt";',
              ],
            },
            outputs: {
              parameters: [
                {
                  name: 'fileLocation',
                  value: '{{inputs.parameters.target-domain}}',
                },
              ],
              artifacts: [
                {
                  name: 'message',
                  path: '/tmp/',
                  globalName: 'my-global-art',
                  archive: { none: {} },
                  s3: {
                    endpoint: 'fra1.digitaloceanspaces.com',
                    bucket: 'testv1',
                    key:
                      '{{workflow.parameters.userId}}/artifacts/{{workflow.name}}/raw/dnsmap/',
                    accessKeySecret: { name: 'artifact-key', key: 'accessKey' },
                    secretKeySecret: { name: 'artifact-key', key: 'secretKey' },
                  },
                },
              ],
            },
          },
        ],
      },
    },
  };
  return startArgo('POST', url, body);
};
