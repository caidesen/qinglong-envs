import type { GenerateServiceProps } from "openapi-ts-request"

export default [
  {
    schemaPath: "../docs/swagger.yaml",
    requestLibPath: "@/utils/http",
    requestImportStatement: `import request, {type RequestOptions} from "@/utils/http"`,
    requestOptionsType: "RequestOptions",
    serversPath: "./src/apis",
  },
] as GenerateServiceProps[]
