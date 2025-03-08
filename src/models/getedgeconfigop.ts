/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import * as z from "zod";
import { safeParse } from "../lib/schemas.js";
import { ClosedEnum } from "../types/enums.js";
import { Result as SafeParseResult } from "../types/fp.js";
import { SDKValidationError } from "./sdkvalidationerror.js";

export type GetEdgeConfigRequest = {
  edgeConfigId: string;
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
 * Keeps track of the current state of the Edge Config while it gets transferred.
 */
export type GetEdgeConfigTransfer = {
  fromAccountId: string;
  startedAt: number;
  doneAt: number | null;
};

export type GetEdgeConfigSchema = {};

export const GetEdgeConfigPurposeEdgeConfigType = {
  Experimentation: "experimentation",
} as const;
export type GetEdgeConfigPurposeEdgeConfigType = ClosedEnum<
  typeof GetEdgeConfigPurposeEdgeConfigType
>;

export type GetEdgeConfigPurpose2 = {
  type: GetEdgeConfigPurposeEdgeConfigType;
  resourceId: string;
};

export const GetEdgeConfigPurposeType = {
  Flags: "flags",
} as const;
export type GetEdgeConfigPurposeType = ClosedEnum<
  typeof GetEdgeConfigPurposeType
>;

export type GetEdgeConfigPurpose1 = {
  type: GetEdgeConfigPurposeType;
  projectId: string;
};

export type GetEdgeConfigPurpose =
  | GetEdgeConfigPurpose1
  | GetEdgeConfigPurpose2;

/**
 * The EdgeConfig.
 */
export type GetEdgeConfigResponseBody = {
  createdAt?: number | undefined;
  updatedAt?: number | undefined;
  id?: string | undefined;
  /**
   * Name for the Edge Config Names are not unique. Must start with an alphabetic character and can contain only alphanumeric characters and underscores).
   */
  slug?: string | undefined;
  ownerId?: string | undefined;
  digest?: string | undefined;
  /**
   * Keeps track of the current state of the Edge Config while it gets transferred.
   */
  transfer?: GetEdgeConfigTransfer | undefined;
  schema?: GetEdgeConfigSchema | undefined;
  purpose?: GetEdgeConfigPurpose1 | GetEdgeConfigPurpose2 | undefined;
  sizeInBytes: number;
  itemCount: number;
};

/** @internal */
export const GetEdgeConfigRequest$inboundSchema: z.ZodType<
  GetEdgeConfigRequest,
  z.ZodTypeDef,
  unknown
> = z.object({
  edgeConfigId: z.string(),
  teamId: z.string().optional(),
  slug: z.string().optional(),
});

/** @internal */
export type GetEdgeConfigRequest$Outbound = {
  edgeConfigId: string;
  teamId?: string | undefined;
  slug?: string | undefined;
};

/** @internal */
export const GetEdgeConfigRequest$outboundSchema: z.ZodType<
  GetEdgeConfigRequest$Outbound,
  z.ZodTypeDef,
  GetEdgeConfigRequest
> = z.object({
  edgeConfigId: z.string(),
  teamId: z.string().optional(),
  slug: z.string().optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetEdgeConfigRequest$ {
  /** @deprecated use `GetEdgeConfigRequest$inboundSchema` instead. */
  export const inboundSchema = GetEdgeConfigRequest$inboundSchema;
  /** @deprecated use `GetEdgeConfigRequest$outboundSchema` instead. */
  export const outboundSchema = GetEdgeConfigRequest$outboundSchema;
  /** @deprecated use `GetEdgeConfigRequest$Outbound` instead. */
  export type Outbound = GetEdgeConfigRequest$Outbound;
}

export function getEdgeConfigRequestToJSON(
  getEdgeConfigRequest: GetEdgeConfigRequest,
): string {
  return JSON.stringify(
    GetEdgeConfigRequest$outboundSchema.parse(getEdgeConfigRequest),
  );
}

export function getEdgeConfigRequestFromJSON(
  jsonString: string,
): SafeParseResult<GetEdgeConfigRequest, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetEdgeConfigRequest$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetEdgeConfigRequest' from JSON`,
  );
}

/** @internal */
export const GetEdgeConfigTransfer$inboundSchema: z.ZodType<
  GetEdgeConfigTransfer,
  z.ZodTypeDef,
  unknown
> = z.object({
  fromAccountId: z.string(),
  startedAt: z.number(),
  doneAt: z.nullable(z.number()),
});

/** @internal */
export type GetEdgeConfigTransfer$Outbound = {
  fromAccountId: string;
  startedAt: number;
  doneAt: number | null;
};

/** @internal */
export const GetEdgeConfigTransfer$outboundSchema: z.ZodType<
  GetEdgeConfigTransfer$Outbound,
  z.ZodTypeDef,
  GetEdgeConfigTransfer
> = z.object({
  fromAccountId: z.string(),
  startedAt: z.number(),
  doneAt: z.nullable(z.number()),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetEdgeConfigTransfer$ {
  /** @deprecated use `GetEdgeConfigTransfer$inboundSchema` instead. */
  export const inboundSchema = GetEdgeConfigTransfer$inboundSchema;
  /** @deprecated use `GetEdgeConfigTransfer$outboundSchema` instead. */
  export const outboundSchema = GetEdgeConfigTransfer$outboundSchema;
  /** @deprecated use `GetEdgeConfigTransfer$Outbound` instead. */
  export type Outbound = GetEdgeConfigTransfer$Outbound;
}

export function getEdgeConfigTransferToJSON(
  getEdgeConfigTransfer: GetEdgeConfigTransfer,
): string {
  return JSON.stringify(
    GetEdgeConfigTransfer$outboundSchema.parse(getEdgeConfigTransfer),
  );
}

export function getEdgeConfigTransferFromJSON(
  jsonString: string,
): SafeParseResult<GetEdgeConfigTransfer, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetEdgeConfigTransfer$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetEdgeConfigTransfer' from JSON`,
  );
}

/** @internal */
export const GetEdgeConfigSchema$inboundSchema: z.ZodType<
  GetEdgeConfigSchema,
  z.ZodTypeDef,
  unknown
> = z.object({});

/** @internal */
export type GetEdgeConfigSchema$Outbound = {};

/** @internal */
export const GetEdgeConfigSchema$outboundSchema: z.ZodType<
  GetEdgeConfigSchema$Outbound,
  z.ZodTypeDef,
  GetEdgeConfigSchema
> = z.object({});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetEdgeConfigSchema$ {
  /** @deprecated use `GetEdgeConfigSchema$inboundSchema` instead. */
  export const inboundSchema = GetEdgeConfigSchema$inboundSchema;
  /** @deprecated use `GetEdgeConfigSchema$outboundSchema` instead. */
  export const outboundSchema = GetEdgeConfigSchema$outboundSchema;
  /** @deprecated use `GetEdgeConfigSchema$Outbound` instead. */
  export type Outbound = GetEdgeConfigSchema$Outbound;
}

export function getEdgeConfigSchemaToJSON(
  getEdgeConfigSchema: GetEdgeConfigSchema,
): string {
  return JSON.stringify(
    GetEdgeConfigSchema$outboundSchema.parse(getEdgeConfigSchema),
  );
}

export function getEdgeConfigSchemaFromJSON(
  jsonString: string,
): SafeParseResult<GetEdgeConfigSchema, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetEdgeConfigSchema$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetEdgeConfigSchema' from JSON`,
  );
}

/** @internal */
export const GetEdgeConfigPurposeEdgeConfigType$inboundSchema: z.ZodNativeEnum<
  typeof GetEdgeConfigPurposeEdgeConfigType
> = z.nativeEnum(GetEdgeConfigPurposeEdgeConfigType);

/** @internal */
export const GetEdgeConfigPurposeEdgeConfigType$outboundSchema: z.ZodNativeEnum<
  typeof GetEdgeConfigPurposeEdgeConfigType
> = GetEdgeConfigPurposeEdgeConfigType$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetEdgeConfigPurposeEdgeConfigType$ {
  /** @deprecated use `GetEdgeConfigPurposeEdgeConfigType$inboundSchema` instead. */
  export const inboundSchema = GetEdgeConfigPurposeEdgeConfigType$inboundSchema;
  /** @deprecated use `GetEdgeConfigPurposeEdgeConfigType$outboundSchema` instead. */
  export const outboundSchema =
    GetEdgeConfigPurposeEdgeConfigType$outboundSchema;
}

/** @internal */
export const GetEdgeConfigPurpose2$inboundSchema: z.ZodType<
  GetEdgeConfigPurpose2,
  z.ZodTypeDef,
  unknown
> = z.object({
  type: GetEdgeConfigPurposeEdgeConfigType$inboundSchema,
  resourceId: z.string(),
});

/** @internal */
export type GetEdgeConfigPurpose2$Outbound = {
  type: string;
  resourceId: string;
};

/** @internal */
export const GetEdgeConfigPurpose2$outboundSchema: z.ZodType<
  GetEdgeConfigPurpose2$Outbound,
  z.ZodTypeDef,
  GetEdgeConfigPurpose2
> = z.object({
  type: GetEdgeConfigPurposeEdgeConfigType$outboundSchema,
  resourceId: z.string(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetEdgeConfigPurpose2$ {
  /** @deprecated use `GetEdgeConfigPurpose2$inboundSchema` instead. */
  export const inboundSchema = GetEdgeConfigPurpose2$inboundSchema;
  /** @deprecated use `GetEdgeConfigPurpose2$outboundSchema` instead. */
  export const outboundSchema = GetEdgeConfigPurpose2$outboundSchema;
  /** @deprecated use `GetEdgeConfigPurpose2$Outbound` instead. */
  export type Outbound = GetEdgeConfigPurpose2$Outbound;
}

export function getEdgeConfigPurpose2ToJSON(
  getEdgeConfigPurpose2: GetEdgeConfigPurpose2,
): string {
  return JSON.stringify(
    GetEdgeConfigPurpose2$outboundSchema.parse(getEdgeConfigPurpose2),
  );
}

export function getEdgeConfigPurpose2FromJSON(
  jsonString: string,
): SafeParseResult<GetEdgeConfigPurpose2, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetEdgeConfigPurpose2$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetEdgeConfigPurpose2' from JSON`,
  );
}

/** @internal */
export const GetEdgeConfigPurposeType$inboundSchema: z.ZodNativeEnum<
  typeof GetEdgeConfigPurposeType
> = z.nativeEnum(GetEdgeConfigPurposeType);

/** @internal */
export const GetEdgeConfigPurposeType$outboundSchema: z.ZodNativeEnum<
  typeof GetEdgeConfigPurposeType
> = GetEdgeConfigPurposeType$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetEdgeConfigPurposeType$ {
  /** @deprecated use `GetEdgeConfigPurposeType$inboundSchema` instead. */
  export const inboundSchema = GetEdgeConfigPurposeType$inboundSchema;
  /** @deprecated use `GetEdgeConfigPurposeType$outboundSchema` instead. */
  export const outboundSchema = GetEdgeConfigPurposeType$outboundSchema;
}

/** @internal */
export const GetEdgeConfigPurpose1$inboundSchema: z.ZodType<
  GetEdgeConfigPurpose1,
  z.ZodTypeDef,
  unknown
> = z.object({
  type: GetEdgeConfigPurposeType$inboundSchema,
  projectId: z.string(),
});

/** @internal */
export type GetEdgeConfigPurpose1$Outbound = {
  type: string;
  projectId: string;
};

/** @internal */
export const GetEdgeConfigPurpose1$outboundSchema: z.ZodType<
  GetEdgeConfigPurpose1$Outbound,
  z.ZodTypeDef,
  GetEdgeConfigPurpose1
> = z.object({
  type: GetEdgeConfigPurposeType$outboundSchema,
  projectId: z.string(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetEdgeConfigPurpose1$ {
  /** @deprecated use `GetEdgeConfigPurpose1$inboundSchema` instead. */
  export const inboundSchema = GetEdgeConfigPurpose1$inboundSchema;
  /** @deprecated use `GetEdgeConfigPurpose1$outboundSchema` instead. */
  export const outboundSchema = GetEdgeConfigPurpose1$outboundSchema;
  /** @deprecated use `GetEdgeConfigPurpose1$Outbound` instead. */
  export type Outbound = GetEdgeConfigPurpose1$Outbound;
}

export function getEdgeConfigPurpose1ToJSON(
  getEdgeConfigPurpose1: GetEdgeConfigPurpose1,
): string {
  return JSON.stringify(
    GetEdgeConfigPurpose1$outboundSchema.parse(getEdgeConfigPurpose1),
  );
}

export function getEdgeConfigPurpose1FromJSON(
  jsonString: string,
): SafeParseResult<GetEdgeConfigPurpose1, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetEdgeConfigPurpose1$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetEdgeConfigPurpose1' from JSON`,
  );
}

/** @internal */
export const GetEdgeConfigPurpose$inboundSchema: z.ZodType<
  GetEdgeConfigPurpose,
  z.ZodTypeDef,
  unknown
> = z.union([
  z.lazy(() => GetEdgeConfigPurpose1$inboundSchema),
  z.lazy(() => GetEdgeConfigPurpose2$inboundSchema),
]);

/** @internal */
export type GetEdgeConfigPurpose$Outbound =
  | GetEdgeConfigPurpose1$Outbound
  | GetEdgeConfigPurpose2$Outbound;

/** @internal */
export const GetEdgeConfigPurpose$outboundSchema: z.ZodType<
  GetEdgeConfigPurpose$Outbound,
  z.ZodTypeDef,
  GetEdgeConfigPurpose
> = z.union([
  z.lazy(() => GetEdgeConfigPurpose1$outboundSchema),
  z.lazy(() => GetEdgeConfigPurpose2$outboundSchema),
]);

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetEdgeConfigPurpose$ {
  /** @deprecated use `GetEdgeConfigPurpose$inboundSchema` instead. */
  export const inboundSchema = GetEdgeConfigPurpose$inboundSchema;
  /** @deprecated use `GetEdgeConfigPurpose$outboundSchema` instead. */
  export const outboundSchema = GetEdgeConfigPurpose$outboundSchema;
  /** @deprecated use `GetEdgeConfigPurpose$Outbound` instead. */
  export type Outbound = GetEdgeConfigPurpose$Outbound;
}

export function getEdgeConfigPurposeToJSON(
  getEdgeConfigPurpose: GetEdgeConfigPurpose,
): string {
  return JSON.stringify(
    GetEdgeConfigPurpose$outboundSchema.parse(getEdgeConfigPurpose),
  );
}

export function getEdgeConfigPurposeFromJSON(
  jsonString: string,
): SafeParseResult<GetEdgeConfigPurpose, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetEdgeConfigPurpose$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetEdgeConfigPurpose' from JSON`,
  );
}

/** @internal */
export const GetEdgeConfigResponseBody$inboundSchema: z.ZodType<
  GetEdgeConfigResponseBody,
  z.ZodTypeDef,
  unknown
> = z.object({
  createdAt: z.number().optional(),
  updatedAt: z.number().optional(),
  id: z.string().optional(),
  slug: z.string().optional(),
  ownerId: z.string().optional(),
  digest: z.string().optional(),
  transfer: z.lazy(() => GetEdgeConfigTransfer$inboundSchema).optional(),
  schema: z.lazy(() => GetEdgeConfigSchema$inboundSchema).optional(),
  purpose: z.union([
    z.lazy(() => GetEdgeConfigPurpose1$inboundSchema),
    z.lazy(() => GetEdgeConfigPurpose2$inboundSchema),
  ]).optional(),
  sizeInBytes: z.number(),
  itemCount: z.number(),
});

/** @internal */
export type GetEdgeConfigResponseBody$Outbound = {
  createdAt?: number | undefined;
  updatedAt?: number | undefined;
  id?: string | undefined;
  slug?: string | undefined;
  ownerId?: string | undefined;
  digest?: string | undefined;
  transfer?: GetEdgeConfigTransfer$Outbound | undefined;
  schema?: GetEdgeConfigSchema$Outbound | undefined;
  purpose?:
    | GetEdgeConfigPurpose1$Outbound
    | GetEdgeConfigPurpose2$Outbound
    | undefined;
  sizeInBytes: number;
  itemCount: number;
};

/** @internal */
export const GetEdgeConfigResponseBody$outboundSchema: z.ZodType<
  GetEdgeConfigResponseBody$Outbound,
  z.ZodTypeDef,
  GetEdgeConfigResponseBody
> = z.object({
  createdAt: z.number().optional(),
  updatedAt: z.number().optional(),
  id: z.string().optional(),
  slug: z.string().optional(),
  ownerId: z.string().optional(),
  digest: z.string().optional(),
  transfer: z.lazy(() => GetEdgeConfigTransfer$outboundSchema).optional(),
  schema: z.lazy(() => GetEdgeConfigSchema$outboundSchema).optional(),
  purpose: z.union([
    z.lazy(() => GetEdgeConfigPurpose1$outboundSchema),
    z.lazy(() => GetEdgeConfigPurpose2$outboundSchema),
  ]).optional(),
  sizeInBytes: z.number(),
  itemCount: z.number(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetEdgeConfigResponseBody$ {
  /** @deprecated use `GetEdgeConfigResponseBody$inboundSchema` instead. */
  export const inboundSchema = GetEdgeConfigResponseBody$inboundSchema;
  /** @deprecated use `GetEdgeConfigResponseBody$outboundSchema` instead. */
  export const outboundSchema = GetEdgeConfigResponseBody$outboundSchema;
  /** @deprecated use `GetEdgeConfigResponseBody$Outbound` instead. */
  export type Outbound = GetEdgeConfigResponseBody$Outbound;
}

export function getEdgeConfigResponseBodyToJSON(
  getEdgeConfigResponseBody: GetEdgeConfigResponseBody,
): string {
  return JSON.stringify(
    GetEdgeConfigResponseBody$outboundSchema.parse(getEdgeConfigResponseBody),
  );
}

export function getEdgeConfigResponseBodyFromJSON(
  jsonString: string,
): SafeParseResult<GetEdgeConfigResponseBody, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetEdgeConfigResponseBody$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetEdgeConfigResponseBody' from JSON`,
  );
}
