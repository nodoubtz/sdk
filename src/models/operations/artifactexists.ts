/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import * as z from "zod";
import { safeParse } from "../../lib/schemas.js";
import { Result as SafeParseResult } from "../../types/fp.js";
import { SDKValidationError } from "../errors/sdkvalidationerror.js";

export type ArtifactExistsRequest = {
  /**
   * The artifact hash
   */
  hash: string;
  /**
   * The Team identifier to perform the request on behalf of.
   */
  teamId?: string | undefined;
  /**
   * The Team slug to perform the request on behalf of.
   */
  slug?: string | undefined;
};

/** @internal */
export const ArtifactExistsRequest$inboundSchema: z.ZodType<
  ArtifactExistsRequest,
  z.ZodTypeDef,
  unknown
> = z.object({
  hash: z.string(),
  teamId: z.string().optional(),
  slug: z.string().optional(),
});

/** @internal */
export type ArtifactExistsRequest$Outbound = {
  hash: string;
  teamId?: string | undefined;
  slug?: string | undefined;
};

/** @internal */
export const ArtifactExistsRequest$outboundSchema: z.ZodType<
  ArtifactExistsRequest$Outbound,
  z.ZodTypeDef,
  ArtifactExistsRequest
> = z.object({
  hash: z.string(),
  teamId: z.string().optional(),
  slug: z.string().optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace ArtifactExistsRequest$ {
  /** @deprecated use `ArtifactExistsRequest$inboundSchema` instead. */
  export const inboundSchema = ArtifactExistsRequest$inboundSchema;
  /** @deprecated use `ArtifactExistsRequest$outboundSchema` instead. */
  export const outboundSchema = ArtifactExistsRequest$outboundSchema;
  /** @deprecated use `ArtifactExistsRequest$Outbound` instead. */
  export type Outbound = ArtifactExistsRequest$Outbound;
}

export function artifactExistsRequestToJSON(
  artifactExistsRequest: ArtifactExistsRequest,
): string {
  return JSON.stringify(
    ArtifactExistsRequest$outboundSchema.parse(artifactExistsRequest),
  );
}

export function artifactExistsRequestFromJSON(
  jsonString: string,
): SafeParseResult<ArtifactExistsRequest, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => ArtifactExistsRequest$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'ArtifactExistsRequest' from JSON`,
  );
}
