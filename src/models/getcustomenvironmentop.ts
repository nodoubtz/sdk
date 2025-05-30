/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import * as z from "zod";
import { safeParse } from "../lib/schemas.js";
import { ClosedEnum } from "../types/enums.js";
import { Result as SafeParseResult } from "../types/fp.js";
import { SDKValidationError } from "./sdkvalidationerror.js";

export type GetCustomEnvironmentRequest = {
  /**
   * The unique project identifier or the project name
   */
  idOrName: string;
  /**
   * The unique custom environment identifier within the project
   */
  environmentSlugOrId: string;
  /**
   * The Team identifier to perform the request on behalf of.
   */
  teamId?: string | undefined;
  /**
   * The Team slug to perform the request on behalf of.
   */
  slug?: string | undefined;
};

/**
 * The type of environment (production, preview, or development)
 */
export const GetCustomEnvironmentType = {
  Production: "production",
  Preview: "preview",
  Development: "development",
} as const;
/**
 * The type of environment (production, preview, or development)
 */
export type GetCustomEnvironmentType = ClosedEnum<
  typeof GetCustomEnvironmentType
>;

/**
 * The type of matching to perform
 */
export const GetCustomEnvironmentEnvironmentType = {
  EndsWith: "endsWith",
  StartsWith: "startsWith",
  Equals: "equals",
} as const;
/**
 * The type of matching to perform
 */
export type GetCustomEnvironmentEnvironmentType = ClosedEnum<
  typeof GetCustomEnvironmentEnvironmentType
>;

/**
 * Configuration for matching git branches to this environment
 */
export type GetCustomEnvironmentBranchMatcher = {
  /**
   * The type of matching to perform
   */
  type: GetCustomEnvironmentEnvironmentType;
  /**
   * The pattern to match against branch names
   */
  pattern: string;
};

/**
 * A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
 */
export type GetCustomEnvironmentVerification = {
  type: string;
  domain: string;
  value: string;
  reason: string;
};

/**
 * List of domains associated with this environment
 */
export type GetCustomEnvironmentDomains = {
  name: string;
  apexName: string;
  projectId: string;
  redirect?: string | null | undefined;
  redirectStatusCode?: number | null | undefined;
  gitBranch?: string | null | undefined;
  customEnvironmentId?: string | null | undefined;
  updatedAt?: number | undefined;
  createdAt?: number | undefined;
  /**
   * `true` if the domain is verified for use with the project. If `false` it will not be used as an alias on this project until the challenge in `verification` is completed.
   */
  verified: boolean;
  /**
   * A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
   */
  verification?: Array<GetCustomEnvironmentVerification> | undefined;
};

/**
 * Internal representation of a custom environment with all required properties
 */
export type GetCustomEnvironmentResponseBody = {
  /**
   * Unique identifier for the custom environment (format: env_*)
   */
  id: string;
  /**
   * URL-friendly name of the environment
   */
  slug: string;
  /**
   * The type of environment (production, preview, or development)
   */
  type: GetCustomEnvironmentType;
  /**
   * Optional description of the environment's purpose
   */
  description?: string | undefined;
  /**
   * Configuration for matching git branches to this environment
   */
  branchMatcher?: GetCustomEnvironmentBranchMatcher | undefined;
  /**
   * List of domains associated with this environment
   */
  domains?: Array<GetCustomEnvironmentDomains> | undefined;
  /**
   * List of aliases for the current deployment
   */
  currentDeploymentAliases?: Array<string> | undefined;
  /**
   * Timestamp when the environment was created
   */
  createdAt: number;
  /**
   * Timestamp when the environment was last updated
   */
  updatedAt: number;
};

/** @internal */
export const GetCustomEnvironmentRequest$inboundSchema: z.ZodType<
  GetCustomEnvironmentRequest,
  z.ZodTypeDef,
  unknown
> = z.object({
  idOrName: z.string(),
  environmentSlugOrId: z.string(),
  teamId: z.string().optional(),
  slug: z.string().optional(),
});

/** @internal */
export type GetCustomEnvironmentRequest$Outbound = {
  idOrName: string;
  environmentSlugOrId: string;
  teamId?: string | undefined;
  slug?: string | undefined;
};

/** @internal */
export const GetCustomEnvironmentRequest$outboundSchema: z.ZodType<
  GetCustomEnvironmentRequest$Outbound,
  z.ZodTypeDef,
  GetCustomEnvironmentRequest
> = z.object({
  idOrName: z.string(),
  environmentSlugOrId: z.string(),
  teamId: z.string().optional(),
  slug: z.string().optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetCustomEnvironmentRequest$ {
  /** @deprecated use `GetCustomEnvironmentRequest$inboundSchema` instead. */
  export const inboundSchema = GetCustomEnvironmentRequest$inboundSchema;
  /** @deprecated use `GetCustomEnvironmentRequest$outboundSchema` instead. */
  export const outboundSchema = GetCustomEnvironmentRequest$outboundSchema;
  /** @deprecated use `GetCustomEnvironmentRequest$Outbound` instead. */
  export type Outbound = GetCustomEnvironmentRequest$Outbound;
}

export function getCustomEnvironmentRequestToJSON(
  getCustomEnvironmentRequest: GetCustomEnvironmentRequest,
): string {
  return JSON.stringify(
    GetCustomEnvironmentRequest$outboundSchema.parse(
      getCustomEnvironmentRequest,
    ),
  );
}

export function getCustomEnvironmentRequestFromJSON(
  jsonString: string,
): SafeParseResult<GetCustomEnvironmentRequest, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetCustomEnvironmentRequest$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetCustomEnvironmentRequest' from JSON`,
  );
}

/** @internal */
export const GetCustomEnvironmentType$inboundSchema: z.ZodNativeEnum<
  typeof GetCustomEnvironmentType
> = z.nativeEnum(GetCustomEnvironmentType);

/** @internal */
export const GetCustomEnvironmentType$outboundSchema: z.ZodNativeEnum<
  typeof GetCustomEnvironmentType
> = GetCustomEnvironmentType$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetCustomEnvironmentType$ {
  /** @deprecated use `GetCustomEnvironmentType$inboundSchema` instead. */
  export const inboundSchema = GetCustomEnvironmentType$inboundSchema;
  /** @deprecated use `GetCustomEnvironmentType$outboundSchema` instead. */
  export const outboundSchema = GetCustomEnvironmentType$outboundSchema;
}

/** @internal */
export const GetCustomEnvironmentEnvironmentType$inboundSchema: z.ZodNativeEnum<
  typeof GetCustomEnvironmentEnvironmentType
> = z.nativeEnum(GetCustomEnvironmentEnvironmentType);

/** @internal */
export const GetCustomEnvironmentEnvironmentType$outboundSchema:
  z.ZodNativeEnum<typeof GetCustomEnvironmentEnvironmentType> =
    GetCustomEnvironmentEnvironmentType$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetCustomEnvironmentEnvironmentType$ {
  /** @deprecated use `GetCustomEnvironmentEnvironmentType$inboundSchema` instead. */
  export const inboundSchema =
    GetCustomEnvironmentEnvironmentType$inboundSchema;
  /** @deprecated use `GetCustomEnvironmentEnvironmentType$outboundSchema` instead. */
  export const outboundSchema =
    GetCustomEnvironmentEnvironmentType$outboundSchema;
}

/** @internal */
export const GetCustomEnvironmentBranchMatcher$inboundSchema: z.ZodType<
  GetCustomEnvironmentBranchMatcher,
  z.ZodTypeDef,
  unknown
> = z.object({
  type: GetCustomEnvironmentEnvironmentType$inboundSchema,
  pattern: z.string(),
});

/** @internal */
export type GetCustomEnvironmentBranchMatcher$Outbound = {
  type: string;
  pattern: string;
};

/** @internal */
export const GetCustomEnvironmentBranchMatcher$outboundSchema: z.ZodType<
  GetCustomEnvironmentBranchMatcher$Outbound,
  z.ZodTypeDef,
  GetCustomEnvironmentBranchMatcher
> = z.object({
  type: GetCustomEnvironmentEnvironmentType$outboundSchema,
  pattern: z.string(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetCustomEnvironmentBranchMatcher$ {
  /** @deprecated use `GetCustomEnvironmentBranchMatcher$inboundSchema` instead. */
  export const inboundSchema = GetCustomEnvironmentBranchMatcher$inboundSchema;
  /** @deprecated use `GetCustomEnvironmentBranchMatcher$outboundSchema` instead. */
  export const outboundSchema =
    GetCustomEnvironmentBranchMatcher$outboundSchema;
  /** @deprecated use `GetCustomEnvironmentBranchMatcher$Outbound` instead. */
  export type Outbound = GetCustomEnvironmentBranchMatcher$Outbound;
}

export function getCustomEnvironmentBranchMatcherToJSON(
  getCustomEnvironmentBranchMatcher: GetCustomEnvironmentBranchMatcher,
): string {
  return JSON.stringify(
    GetCustomEnvironmentBranchMatcher$outboundSchema.parse(
      getCustomEnvironmentBranchMatcher,
    ),
  );
}

export function getCustomEnvironmentBranchMatcherFromJSON(
  jsonString: string,
): SafeParseResult<GetCustomEnvironmentBranchMatcher, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetCustomEnvironmentBranchMatcher$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetCustomEnvironmentBranchMatcher' from JSON`,
  );
}

/** @internal */
export const GetCustomEnvironmentVerification$inboundSchema: z.ZodType<
  GetCustomEnvironmentVerification,
  z.ZodTypeDef,
  unknown
> = z.object({
  type: z.string(),
  domain: z.string(),
  value: z.string(),
  reason: z.string(),
});

/** @internal */
export type GetCustomEnvironmentVerification$Outbound = {
  type: string;
  domain: string;
  value: string;
  reason: string;
};

/** @internal */
export const GetCustomEnvironmentVerification$outboundSchema: z.ZodType<
  GetCustomEnvironmentVerification$Outbound,
  z.ZodTypeDef,
  GetCustomEnvironmentVerification
> = z.object({
  type: z.string(),
  domain: z.string(),
  value: z.string(),
  reason: z.string(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetCustomEnvironmentVerification$ {
  /** @deprecated use `GetCustomEnvironmentVerification$inboundSchema` instead. */
  export const inboundSchema = GetCustomEnvironmentVerification$inboundSchema;
  /** @deprecated use `GetCustomEnvironmentVerification$outboundSchema` instead. */
  export const outboundSchema = GetCustomEnvironmentVerification$outboundSchema;
  /** @deprecated use `GetCustomEnvironmentVerification$Outbound` instead. */
  export type Outbound = GetCustomEnvironmentVerification$Outbound;
}

export function getCustomEnvironmentVerificationToJSON(
  getCustomEnvironmentVerification: GetCustomEnvironmentVerification,
): string {
  return JSON.stringify(
    GetCustomEnvironmentVerification$outboundSchema.parse(
      getCustomEnvironmentVerification,
    ),
  );
}

export function getCustomEnvironmentVerificationFromJSON(
  jsonString: string,
): SafeParseResult<GetCustomEnvironmentVerification, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetCustomEnvironmentVerification$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetCustomEnvironmentVerification' from JSON`,
  );
}

/** @internal */
export const GetCustomEnvironmentDomains$inboundSchema: z.ZodType<
  GetCustomEnvironmentDomains,
  z.ZodTypeDef,
  unknown
> = z.object({
  name: z.string(),
  apexName: z.string(),
  projectId: z.string(),
  redirect: z.nullable(z.string()).optional(),
  redirectStatusCode: z.nullable(z.number()).optional(),
  gitBranch: z.nullable(z.string()).optional(),
  customEnvironmentId: z.nullable(z.string()).optional(),
  updatedAt: z.number().optional(),
  createdAt: z.number().optional(),
  verified: z.boolean(),
  verification: z.array(
    z.lazy(() => GetCustomEnvironmentVerification$inboundSchema),
  ).optional(),
});

/** @internal */
export type GetCustomEnvironmentDomains$Outbound = {
  name: string;
  apexName: string;
  projectId: string;
  redirect?: string | null | undefined;
  redirectStatusCode?: number | null | undefined;
  gitBranch?: string | null | undefined;
  customEnvironmentId?: string | null | undefined;
  updatedAt?: number | undefined;
  createdAt?: number | undefined;
  verified: boolean;
  verification?: Array<GetCustomEnvironmentVerification$Outbound> | undefined;
};

/** @internal */
export const GetCustomEnvironmentDomains$outboundSchema: z.ZodType<
  GetCustomEnvironmentDomains$Outbound,
  z.ZodTypeDef,
  GetCustomEnvironmentDomains
> = z.object({
  name: z.string(),
  apexName: z.string(),
  projectId: z.string(),
  redirect: z.nullable(z.string()).optional(),
  redirectStatusCode: z.nullable(z.number()).optional(),
  gitBranch: z.nullable(z.string()).optional(),
  customEnvironmentId: z.nullable(z.string()).optional(),
  updatedAt: z.number().optional(),
  createdAt: z.number().optional(),
  verified: z.boolean(),
  verification: z.array(
    z.lazy(() => GetCustomEnvironmentVerification$outboundSchema),
  ).optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetCustomEnvironmentDomains$ {
  /** @deprecated use `GetCustomEnvironmentDomains$inboundSchema` instead. */
  export const inboundSchema = GetCustomEnvironmentDomains$inboundSchema;
  /** @deprecated use `GetCustomEnvironmentDomains$outboundSchema` instead. */
  export const outboundSchema = GetCustomEnvironmentDomains$outboundSchema;
  /** @deprecated use `GetCustomEnvironmentDomains$Outbound` instead. */
  export type Outbound = GetCustomEnvironmentDomains$Outbound;
}

export function getCustomEnvironmentDomainsToJSON(
  getCustomEnvironmentDomains: GetCustomEnvironmentDomains,
): string {
  return JSON.stringify(
    GetCustomEnvironmentDomains$outboundSchema.parse(
      getCustomEnvironmentDomains,
    ),
  );
}

export function getCustomEnvironmentDomainsFromJSON(
  jsonString: string,
): SafeParseResult<GetCustomEnvironmentDomains, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetCustomEnvironmentDomains$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetCustomEnvironmentDomains' from JSON`,
  );
}

/** @internal */
export const GetCustomEnvironmentResponseBody$inboundSchema: z.ZodType<
  GetCustomEnvironmentResponseBody,
  z.ZodTypeDef,
  unknown
> = z.object({
  id: z.string(),
  slug: z.string(),
  type: GetCustomEnvironmentType$inboundSchema,
  description: z.string().optional(),
  branchMatcher: z.lazy(() => GetCustomEnvironmentBranchMatcher$inboundSchema)
    .optional(),
  domains: z.array(z.lazy(() => GetCustomEnvironmentDomains$inboundSchema))
    .optional(),
  currentDeploymentAliases: z.array(z.string()).optional(),
  createdAt: z.number(),
  updatedAt: z.number(),
});

/** @internal */
export type GetCustomEnvironmentResponseBody$Outbound = {
  id: string;
  slug: string;
  type: string;
  description?: string | undefined;
  branchMatcher?: GetCustomEnvironmentBranchMatcher$Outbound | undefined;
  domains?: Array<GetCustomEnvironmentDomains$Outbound> | undefined;
  currentDeploymentAliases?: Array<string> | undefined;
  createdAt: number;
  updatedAt: number;
};

/** @internal */
export const GetCustomEnvironmentResponseBody$outboundSchema: z.ZodType<
  GetCustomEnvironmentResponseBody$Outbound,
  z.ZodTypeDef,
  GetCustomEnvironmentResponseBody
> = z.object({
  id: z.string(),
  slug: z.string(),
  type: GetCustomEnvironmentType$outboundSchema,
  description: z.string().optional(),
  branchMatcher: z.lazy(() => GetCustomEnvironmentBranchMatcher$outboundSchema)
    .optional(),
  domains: z.array(z.lazy(() => GetCustomEnvironmentDomains$outboundSchema))
    .optional(),
  currentDeploymentAliases: z.array(z.string()).optional(),
  createdAt: z.number(),
  updatedAt: z.number(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetCustomEnvironmentResponseBody$ {
  /** @deprecated use `GetCustomEnvironmentResponseBody$inboundSchema` instead. */
  export const inboundSchema = GetCustomEnvironmentResponseBody$inboundSchema;
  /** @deprecated use `GetCustomEnvironmentResponseBody$outboundSchema` instead. */
  export const outboundSchema = GetCustomEnvironmentResponseBody$outboundSchema;
  /** @deprecated use `GetCustomEnvironmentResponseBody$Outbound` instead. */
  export type Outbound = GetCustomEnvironmentResponseBody$Outbound;
}

export function getCustomEnvironmentResponseBodyToJSON(
  getCustomEnvironmentResponseBody: GetCustomEnvironmentResponseBody,
): string {
  return JSON.stringify(
    GetCustomEnvironmentResponseBody$outboundSchema.parse(
      getCustomEnvironmentResponseBody,
    ),
  );
}

export function getCustomEnvironmentResponseBodyFromJSON(
  jsonString: string,
): SafeParseResult<GetCustomEnvironmentResponseBody, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetCustomEnvironmentResponseBody$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetCustomEnvironmentResponseBody' from JSON`,
  );
}
