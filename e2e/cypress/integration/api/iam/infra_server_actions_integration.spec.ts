let withInfraServersListActionToken = '';
let withoutInfraServersListActionToken = '';

const cypressPrefix = 'infra-server-actions';
const policyId1 = `${cypressPrefix}-pol-1-${Cypress.moment().format('MMDDYYhhmm')}`;
const policyId2 = `${cypressPrefix}-pol-2-${Cypress.moment().format('MMDDYYhhmm')}`;
const tokenId1 = `${cypressPrefix}-token-1-${Cypress.moment().format('MMDDYYhhmm')}`;
const tokenId2 = `${cypressPrefix}-token-2-${Cypress.moment().format('MMDDYYhhmm')}`;
const objectsToCleanUp = ['tokens', 'policies'];

const withInfraServersListPolicy = {
  id: policyId1,
  name: tokenId1,
  projects: [],
  members: [`token:${tokenId1}`],
  statements: [
    {
        effect: "ALLOW",
        actions: [
            "infra:infraServers:list",
            "infra:nodes:list"
        ],
        projects: ["*"]
    }]
};


const withoutInfraServersListPolicy = {
    id: policyId2,
    name: tokenId2,
    projects: [],
    members: [`token:${tokenId2}`],
    statements: [
      {
          effect: "DENY",
          actions: [
            "infra:infraServers:list",
            "infra:nodes:list"
          ],
          projects: ["*"]
      }]
  };

// these tests are best read sequentially, as they share state
describe('Infra servers list api', () => {
    before(() => {
        // TODO cleanup projects in before block (can't do now bc we have a project
        // limit and cereal runs async to delete policies)
        cy.cleanupIAMObjectsByIDPrefixes(cypressPrefix, objectsToCleanUp);

        cy.request({
            headers: { 'api-token': Cypress.env('ADMIN_TOKEN') },
            method: 'POST',
            url: '/apis/iam/v2/tokens',
            body: {
              id: tokenId1,
              name: tokenId1
            }
          }).then((resp) => {
            withInfraServersListActionToken = resp.body.token.value;
          });

        cy.request({
            headers: { 'api-token': Cypress.env('ADMIN_TOKEN') },
            method: 'POST',
            url: '/apis/iam/v2/policies',
            failOnStatusCode: false,
            body: withInfraServersListPolicy
          }).then((resp) => {
            expect(resp.status).to.equal(200);
          });

        cy.request({
            headers: { 'api-token': Cypress.env('ADMIN_TOKEN') },
            method: 'POST',
            url: '/apis/iam/v2/tokens',
            body: {
              id: tokenId2,
              name: tokenId2
            }
          }).then((resp) => {
            withoutInfraServersListActionToken = resp.body.token.value;
          });


        cy.request({
            headers: { 'api-token': Cypress.env('ADMIN_TOKEN') },
            method: 'POST',
            url: '/apis/iam/v2/policies',
            failOnStatusCode: false,
            body: withoutInfraServersListPolicy
          }).then((resp) => {
            expect(resp.status).to.equal(200);
          });
        })
    }) 
    
    after(() => {
        cy.cleanupIAMObjectsByIDPrefixes(cypressPrefix, objectsToCleanUp);
      });

    it(`infra servers list returns 200 when infraServers list actions is allowed`, () => {
    cy.request({
        headers: { 'api-token': withInfraServersListActionToken, 'content-type': 'application/json+lax' },
        method: 'GET',
        url: `/api/v0/infra/servers`
    }).then((resp) => {
        assert.equal(resp.status, 200);
        });
    }); 

    it(`infra servers list returns 401 when infraServers list actions is deneyed`, () => {
        cy.request({
            headers: { 'api-token': withoutInfraServersListActionToken, 'content-type': 'application/json+lax' },
            method: 'GET',
            url: `/api/v0/infra/servers`
        }).then((resp) => {
            assert.equal(resp.status, 401);
        });
    }); 