Generate a GitHub Actions workflow skeleton from this mermaid diagram.
 
This workflow should:
  * run on ubuntu-latest only
  * use matrix to handle multiple nodes versions
  * use https://registry.npmjs.org as the registry URL
  * use NPM_TOKEN Actions secret for authentication


```mermaid
graph TD;
    Checkout_repository-->Lint_Code;
    Lint_Code-->Setup_Node_12;
    Lint_Code-->Setup_Node_14;
    Lint_Code-->Setup_Node_16;
    Setup_Node_12-->Npm_Ci;
    Setup_Node_14-->Npm_Ci;
    Setup_Node_16-->Npm_Ci;
    Npm_Ci-->Run_Tests;
    Run_Tests-->Compute_Coverage;
    Run_Tests-->Build;
    Build-->Publish;
```