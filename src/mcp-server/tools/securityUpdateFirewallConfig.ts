/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { securityUpdateFirewallConfig } from "../../funcs/securityUpdateFirewallConfig.js";
import { UpdateFirewallConfigRequest$inboundSchema } from "../../models/updatefirewallconfigop.js";
import { formatResult, ToolDefinition } from "../tools.js";

const args = {
  request: UpdateFirewallConfigRequest$inboundSchema,
};

export const tool$securityUpdateFirewallConfig: ToolDefinition<typeof args> = {
  name: "security-update-firewall-config",
  description: `Update Firewall Configuration

Process updates to modify the existing firewall config for a project`,
  args,
  tool: async (client, args, ctx) => {
    const [result, apiCall] = await securityUpdateFirewallConfig(
      client,
      args.request,
      { fetchOptions: { signal: ctx.signal } },
    ).$inspect();

    if (!result.ok) {
      return {
        content: [{ type: "text", text: result.error.message }],
        isError: true,
      };
    }

    const value = result.value;

    return formatResult(value, apiCall);
  },
};
