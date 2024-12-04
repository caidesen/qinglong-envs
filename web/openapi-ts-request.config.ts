import type { GenerateServiceProps } from "openapi-ts-request"

export default [
  {
    schemaPath: "../docs/swagger.yaml",
    requestLibPath: "@/utils/http",
    // hook: {
    //   customFunctionName(data: APIDataType) {
    //     return ''
    //   }
    // },
  },
] as GenerateServiceProps[]
