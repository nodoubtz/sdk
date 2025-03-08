/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { expect, test } from "vitest";
import { Vercel } from "../index.js";
import { createTestHTTPClient } from "./testclient.js";

test("Checks Create Check", async () => {
  const testHttpClient = createTestHTTPClient("createCheck");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.checks.createCheck({
    deploymentId: "dpl_2qn7PZrx89yxY34vEZPD31Y9XVj6",
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
    requestBody: {
      name: "Performance Check",
      path: "/",
      blocking: true,
      detailsUrl: "http://example.com",
      externalId: "1234abc",
      rerequestable: true,
    },
  });
  expect(result).toBeDefined();
  expect(result).toEqual({
    id: "chk_1a2b3c4d5e6f7g8h9i0j",
    name: "Performance Check",
    status: "completed",
    blocking: false,
    integrationId: "<id>",
    deploymentId: "<id>",
    createdAt: 2396.37,
    updatedAt: 1065.29,
  });
});

test("Checks Get All Checks", async () => {
  const testHttpClient = createTestHTTPClient("getAllChecks");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.checks.getAllChecks({
    deploymentId: "dpl_2qn7PZrx89yxY34vEZPD31Y9XVj6",
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
  expect(result).toBeDefined();
  expect(result).toEqual({
    checks: [
      {
        createdAt: 4203.11,
        id: "<id>",
        integrationId: "<id>",
        name: "<value>",
        rerequestable: false,
        status: "registered",
        updatedAt: 4461.16,
      },
    ],
  });
});

test("Checks Get Check", async () => {
  const testHttpClient = createTestHTTPClient("getCheck");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.checks.getCheck({
    deploymentId: "dpl_2qn7PZrx89yxY34vEZPD31Y9XVj6",
    checkId: "check_2qn7PZrx89yxY34vEZPD31Y9XVj6",
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
  expect(result).toBeDefined();
  expect(result).toEqual({
    id: "<id>",
    name: "<value>",
    status: "completed",
    blocking: true,
    integrationId: "<id>",
    deploymentId: "<id>",
    createdAt: 2039.14,
    updatedAt: 676.34,
  });
});

test("Checks Update Check", async () => {
  const testHttpClient = createTestHTTPClient("updateCheck");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.checks.updateCheck({
    deploymentId: "dpl_2qn7PZrx89yxY34vEZPD31Y9XVj6",
    checkId: "check_2qn7PZrx89yxY34vEZPD31Y9XVj6",
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
    requestBody: {
      name: "Performance Check",
      path: "/",
      detailsUrl: "https://example.com/check/run/1234abc",
      output: {
        metrics: {
          fcp: {
            value: 1200,
            previousValue: 900,
            source: "web-vitals",
          },
          lcp: {
            value: 1200,
            previousValue: 1000,
            source: "web-vitals",
          },
          cls: {
            value: 4,
            previousValue: 2,
            source: "web-vitals",
          },
          tbt: {
            value: 3000,
            previousValue: 3500,
            source: "web-vitals",
          },
          virtualExperienceScore: {
            value: 30,
            previousValue: 35,
            source: "web-vitals",
          },
        },
      },
      externalId: "1234abc",
    },
  });
  expect(result).toBeDefined();
  expect(result).toEqual({
    id: "<id>",
    name: "<value>",
    status: "completed",
    blocking: true,
    integrationId: "<id>",
    deploymentId: "<id>",
    createdAt: 9017.64,
    updatedAt: 7909.85,
  });
});

test("Checks Rerequest Check", async () => {
  const testHttpClient = createTestHTTPClient("rerequestCheck");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.checks.rerequestCheck({
    deploymentId: "dpl_2qn7PZrx89yxY34vEZPD31Y9XVj6",
    checkId: "check_2qn7PZrx89yxY34vEZPD31Y9XVj6",
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
  expect(result).toBeDefined();
  expect(result).toEqual({});
});
