import { requestHeaders } from './apiauth';
import { API } from './types';
import { ensureSetting } from './ensure';

export enum Policy {
  Label = 'label',
}

export function resetPolicy(api: API, policy: Policy) {
  cy.request({
    method: 'DELETE',
    url: `${api.mgmtBaseURL}/policies/${policy}`,
    headers: requestHeaders(api),
  }).then((res) => {
    expect(res.status).to.equal(200);
    return null;
  });
}

export function ensureDomainPolicy(
  api: API,
  userLoginMustBeDomain: boolean,
  validateOrgDomains: boolean,
  smtpSenderAddressMatchesInstanceDomain: boolean,
): Cypress.Chainable<null> {
  return cy
    .request({
      method: 'PUT',
      url: `${api.adminBaseURL}/policies/domain`,
      body: {
        userLoginMustBeDomain: userLoginMustBeDomain,
        validateOrgDomains: validateOrgDomains,
        smtpSenderAddressMatchesInstanceDomain: smtpSenderAddressMatchesInstanceDomain,
      },
      failOnStatusCode: false,
      ...auth(api),
    })
    .then((res) => {
      if (!res.isOkStatusCode) {
        expect(res.status).to.equal(400);
        expect(res.body.message).to.contain('Org IAM Policy has not been changed');
      }
      return null;
    });
}
/*
export function legacyEnsureDomainPolicy(
  api: API,
  userLoginMustBeDomain: boolean,
  validateOrgDomains: boolean,
  smtpSenderAddressMatchesInstanceDomain: boolean,
): Cypress.Chainable<number> {
  return ensureSetting(
    api,
    `${api.adminBaseURL}/policies/domain`,
    (body: any) => {
      const result = {
        sequence: parseInt(<string>body.policy?.details?.sequence),
        id: body.policy?.details?.resourceOwner,
        entity: null,
      };
      if (
        body.policy &&
        (body.policy.userLoginMustBeDomain ? body.policy.userLoginMustBeDomain : false) == userLoginMustBeDomain &&
        (body.policy.validateOrgDomains ? body.policy.validateOrgDomains : false) == validateOrgDomains &&
        (body.policy.smtpSenderAddressMatchesInstanceDomain ? body.policy.smtpSenderAddressMatchesInstanceDomain : false) ==
          smtpSenderAddressMatchesInstanceDomain
      ) {
        return { ...result, entity: body.policy };
      }
      return result;
    },
    `${api.adminBaseURL}/policies/domain`,
    {
      userLoginMustBeDomain: userLoginMustBeDomain,
      validateOrgDomains: validateOrgDomains,
      smtpSenderAddressMatchesInstanceDomain: smtpSenderAddressMatchesInstanceDomain,
    },
  );
}
*/
function auth(api: API) {
  return { auth: { bearer: api.token } };
}
