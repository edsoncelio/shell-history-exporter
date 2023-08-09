tf apply --auto-approve
tf apply --auto-approve
tf apply --auto-approve
tf destroy
k get pod -n karpenter
k delete pod karpenter-6c5c6d7f97-gfmtz karpenter-6c5c6d7f97-n2jfp -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

kubectl get pods -n kube-system -l k8s-app=aws-node -o wide
aws ec2 describe-subnets --filters "Name=vpc-id,Values= VPCID" | jq '.Subnets[] | .SubnetId + "=" + "\(.AvailableIpAddressCount)"'\

aws ec2 describe-subnets --filters "Name=vpc-us,Values=vpc-01472c44ec8cf0342" | jq '.Subnets[] | .SubnetId + "=" + "\(.AvailableIpAddressCount)"'\

kubectl label namespace prod-v3 istio-injection-
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k get pod -A grep | Terminating
k get pod -A | grep Terminating
for p in $(kubectl get pods -n prod-v3 | grep Terminating | awk '{print $1}'); do kubectl delete pod  -n prod-v3 $p --grace-period=0 --force;done\

kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

aws eks describe-addon --cluster-name my-cluster --addon-name vpc-cni --query addon.addonVersion --output text
aws eks describe-addon --cluster-name ck_new --addon-name vpc-cni --query addon.addonVersion --output text
kubectl describe daemonset aws-node --namespace kube-system | grep amazon-k8s-cni: | cut -d : -f 3\

kubectl describe daemonset aws-node -n kube-system | grep Image | cut -d "/" -f 2\

k get node -l environment=contraktor-prod
k get provisioner
k delete provisioner contraktor-prod
k get node -l environment=cksign-prod
aws sso login --profile contraktor
aws sso login --profile contraktor
tf init
tf plan
tf plan
tf apply
tf apply
aws eks update-kubeconfig --region us-east-1 --name contraktor-staging
k get pod -A
ls
mkdir apps
k create namespace cksign-stg
k apply -f apps/ck-ai
k get all -n cksign-stg
k get all -n cksign-stg
k apply -f apps/ck-ai
k delete -f apps/ck-ai
k apply -f apps/ck-ai
nslookup ai-dev-new.contraktor.com.br
curl -I https://ai-dev-new.contraktor.com.br/
curl -I https://ai-dev-new.contraktor.com.br/
nslookup ai-dev-new.contraktor.com.br
k delete -f apps/ck-ai
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k apply -f apps/ck-ai
k apply -f apps/ck-ai
k delete -f apps/ck-ai
k apply -f apps/ck-ai
k apply -f apps/ck-ai
curl -I k8s-stagingingress-1fa24aae5e-1271421936.us-east-1.elb.amazonaws.com
k get pod -n karpenter
k delete pod karpenter-6b8df4f4f4-m6vpq karpenter-6b8df4f4f4-xqvh8 -n karpenter
curl -I https://ai-dev-new.contraktor.com.br/
curl -I ai-dev-new.contraktor.com.br
curl https://ai-dev-new.contraktor.com.br/
k apply -f apps/ck-ai
curl -I ai-dev-new.contraktor.com.br
curl -I https://ai-dev-new.contraktor.com.br
k delete -f apps/ck-ai
k apply -f apps/ck-ai
k delete -f apps/ck-ai
k get provisioner
k describe provisioner default
k apply -f apps/ck-ai
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
tf apply 
k get cm aws-auth -n kube-system -o yaml
tf apply 
k get pod -n karpenter
k delete pod karpenter-6b8df4f4f4-lpkrs karpenter-6b8df4f4f4-nh9wd -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
tf apply --auto-approve
tf destroy --auto-approve
tf destroy --auto-approve
tf destroy --auto-approve
tf state rm kubectl_manifest.karpenter_node_template
tf state rm kubectl_manifest.karpenter_provisioner
tf destroy --auto-approve
aws sso login --profile contraktor
tf apply
aws eks update-kubeconfig --region us-east-1 --name contraktor-staging
k get pod -A
k create namespace cksign-stg
k apply -f apps
k apply -f apps/ck-ai
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k get cm aws-auth -n kube-system -o yaml
k delete -f apps/ck-ai
k apply -f apps/ck-ai
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k delete -f apps/ck-ai
kubectl delete node -l karpenter.sh/provisioner-name=default
tf destroy 
ls
tree -L 2
ls
mv ../../../../qa-deployments/dev/starter/ apps
ls
ls apps
cd apps
ls
mkdir cksign-stg
ls
code .
ls
mv starter/ cksign-stg
aws sso login --profile contraktor
aws sso login --profile contraktor
k get node -l infra=starter
aws eks update-kubeconfig --region us-east-1 --name ck_ne
aws eks update-kubeconfig --region us-east-1 --name ck_new
k get node -l infra=starter
k top node -l infra=starter
k describe node ip-10-6-56-28.ec2.internal 
k get node -l infra=starter
k describe node ip-10-6-56-207.ec2.internal
k get node -l infra=starter
k top node -l infra=starter
kubectl label nodes ip-10-6-50-16.ec2.internal infra=starter_api
kubectl label nodes ip-10-6-50-16.ec2.internal environment=starter_api
k top node ip-10-6-50-16.ec2.internal
aws sso login --profile contraktor
aws sso login --profile contraktor
aws sso login --profile contraktor
cd ..
code .
tf init
tf plan
tf plan
tf apply
tf destroy
tf apply
code ../ck_new/karpenter
tf apply
ls
mkdir infra
mv *.tf infra
ls
mv files infra
ls
cd infra
tf plan
tf init
tf plan
tf apply
aws eks update-kubeconfig --name ck-staging --region us-east-1\

ls
cd ..
ls
cd apps
ls
cd cksign-
rm -rf cksign-prd
ls
k apply -f cksign-stg/ck-ai
k create namespace ck-ai-stg
k get pod -A
k apply -f cksign-stg/ck-ai
k apply -k cksign-stg/ck-ai
k get all -n ck-ai-stg
k describe pod -n ck-ai-stg
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k get nodes
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k get pod -n karpenter
k delete pod karpenter-6b8df4f4f4-g6r8q karpenter-6b8df4f4f4-n2f7s
k delete pod karpenter-6b8df4f4f4-g6r8q karpenter-6b8df4f4f4-n2f7s -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k get node
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k get pod -A
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

aws sts decode-authorization-message --encoded-message qFS4GT1Ys3PSdfGR4Q8IkCQjOJZDt8yLmyX8b6uNMaCqt-OU3r6eWJWfFl5vFVIVDZmxxyUuJdcxFvO4kdZdM6V_KI-w-NO4absyyHbDLIHoFiqZFsHZ-ufIQwqIDVgxionHUrLjE0FS58P0ew4O8unVEJcEllFg8DRd243gl-U1G2-ZqUGy_koWduZ6MhZ51gAqC9kpHN7YLTWO3_lhocTUHaXBgbPDWK3inUump-y4pQws7QQIXAr5sG-SU27tSdFzqD0wH_ndcSWo2hdhh-ZBFt87rIE4A-7Brzwxt8Or7P1v_Mrn5pOUqeat_5pF84IbhnIIfeHjixQe3N_D5gYIMQwppbZNFNegbXdRswJfQmY6s0HI1ud6jlYi5hXSg-HYHG76QqEvKrAIoKPMm9kFPrFalbG5i-rsfq-xoV40KZV8A_VOgqEj0GlnlOir1l4Hqwvyb9pH0NOlrtKbpOD8dYryMrWcleLp0gVGYcchyUDCw6CHWvFvyDqqP6SFf8WdlRwZ8X8vAUE98D3VKJUwOTrGn5XvFosxk9_a3U6fbKCjJrdxgbExcdVgIHQek2SYSayz1vSgnAu3ZdopPnyppT6Vu85MVmGq1G2_zyqTsPJxjF_KpThmzBM2TbaN9npmk6zaEcks_eYjuJGx6c50FWBV05tmTGqr_zfKJBWol0uBxOKIg-EmI2NOag6xke4aiHHvemFB5CwkJRHx-10sB8GdvPZ4grUVzcMee6YKr0fzEHidarjfgzqUHkl0_fs2VFHaZHOvmxwa1pmyfT8N
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

cd ..
cd infra
tf apply --auto-approve
k get pod -n karpenter
k get pod -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

tf destroy --auto-approve
aws sso login --profile contraktor
tf destroy
kubectx
kubectx arn:aws:eks:us-east-1:427679376461:cluster/ck_new
k top node arn:aws:eks:us-east-1:427679376461:cluster/ck_new
k top node ip-10-6-50-16.ec2.internal\

ls
ls
tf apply --auto-approve
tf destroy
tf state rm module.eks.kubernetes_config_map_v1_data.aws_auth[0]
tf state rm module.eks.kubernetes_config_map_v1_data.aws_auth
tf state rm kubectl_manifest.test
tf state rm helm_release.karpenter
tf state rm kubernetes_service_account.service-account
tf destroy
tf state rm kubectl_manifest.karpenter_node_template
tm state rm kubectl_manifest.karpenter_provisioner
tf state rm kubectl_manifest.karpenter_provisioner
tf state rm helm_release.lb
tf destroy
tf apply
tf apply
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k get pod -n karpenter
k delete pod karpenter-6c5c6d7f97-hwx5s karpenter-6c5c6d7f97-rlt96 -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
kubectx arn:aws:eks:us-east-1:427679376461:cluster/ck-staging
k get pod -A
aws eks update-kubeconfig --name ck-staging --region us-east-1\

k get pod -A
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k get pod -n karpenter
k delete pod karpenter-6b8df4f4f4-977ck karpenter-6b8df4f4f4-hvzjk -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k create deployment --image=nginx 
k create deployment nginx --image=nginx 
k get pod
k scale deploy nginx --replicas=3
k get pod
k get pod
k scale deploy nginx --replicas=10
k get pod
k get pod
k scale deploy nginx --replicas=40
k get deploy
k get deploy
k get deploy
k get pod
k get deploy
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
tf apply --auto-approve
k get pod -n karpenter
k delete pod karpenter-6b8df4f4f4-kslhz karpenter-6b8df4f4f4-tqdm5 -n karpenter
k get pod -n karpenter
k get pod -n karpenter
k scale deploy nginx --replicas=0
k get pod -n karpenter
k get deploy
k get pod -n karpenter
k get pod -n karpenter
k delete pod karpenter-6b8df4f4f4-kslhz karpenter-6b8df4f4f4-tqdm5 -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k scale deploy nginx --replicas=40
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k get pod -n karpenter
k get deploy
k scale deploy nginx --replicas=50
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k scale deploy nginx --replicas=0
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k get nodes
k get nodes
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k describe provisioner
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
tf apply --auto-approve
k get nodes
k get pod -n karpenter
k delete pod karpenter-6b8df4f4f4-prd89 karpenter-6b8df4f4f4-rkglb -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
helm repo update
tf apply
helm list -A
helm uninstall karpenter -n karpenter
tf apply
tf apply
k get pod -n karpenter
helm list -A
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k get node
k scale deploy nginx --replicas=40
k get nodes
k get nodes
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k get nodes
k get nodes
k scale deploy nginx --replicas=0
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k get nodes
k get nodes
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k get pod
tf apply --auto-approve
k get pod -n karpenter
k delete pod karpenter-6dcdcdb889-lkrqw karpenter-6dcdcdb889-mc274 -n karpenter
k get pod -n karpenter
k get node
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k scale deploy nginx --replicas=30
k get deploy
k get deploy
k get deploy
k get deploy
k get deploy
k get deploy
k get pod
k get deploy
k get deploy
k get deploy
k get deploy
k get deploy
k get deploy
k scale deploy nginx --replicas=0
k get node
k get node
k get node
k get node
code ../../ck_new/karpenter
tf apply
tf apply
k get pod -n karpenter
k delete pod karpenter-6dcdcdb889-bxptg karpenter-6dcdcdb889-rj522 -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
ls
cd ..
ls
cd apps
ls
k apply -k cksign-stg/ck-ai
k get deploy -n cksign-st
k get deploy -n ck-ai-stg
k get deploy -n ck-ai-stg
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k get deploy -n ck-ai-stg
k get deploy -n ck-ai-stg
k get pod -n ck-ai-stg
k get pod -n ck-ai-stg
k get pod -n ck-ai-stg
k describe pod ck-ai-7c576bb955-d6vhx -n ck-ai-stg
k get pod -n ck-ai-stg
k apply -k cksign-stg/starter
k apply -k cksign-stg/starter/
k apply -k cksign-stg/starter/
k apply -k cksign-stg/starter/
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k get node
k get pod -n dev-starter
k apply -k cksign-stg/starter/
k get deploy -n dev-starter
k get pod -n dev-starter
k delete deploy -n dev-starter-superadmin
k delete deploy dev-starter-superadmin -n dev-starter
k delete deploy -n dev-starter-superadmin
k get pod -n dev-starter
k get pod -n dev-starter
k get pod -n dev-starter
k describe pod dev-starter-ck-pay-76f96b8dc8-ndzv6 -n dev-starter
k get pod -n dev-starter
cat ~/.kube/config
k apply -k cksign-stg/starter/
kubectl get deployment -n kube-system aws-load-balancer-controller -o yaml | grep ck_staging 
kubectl get deployment -n kube-system aws-load-balancer-controller -o yaml | grep "ck_staging" 
kubectl get deployment -n kube-system aws-load-balancer-controller -o yaml
kubectl get deployment -n kube-system aws-load-balancer-controller -o yaml | grep ck-staging
k delete -k cksign-stg/starter/
k apply -k cksign-stg/starter/
k apply -k cksign-stg/starter/
k delete -k cksign-stg/ck-ai
k apply -k cksign-stg/ck-ai
k delete -k cksign-stg/ck-ai
k apply -k cksign-stg/ck-ai
cd ..
ls
git status
ls
git status
cd infra
rm -rf .terraform
ls
rm errored.tfstate
ls
cd ..
git status
cd ..
git status
git add contraktor-staging
git rm --cached resources/eks-clusters/contraktor-staging/.terraform/
cd contraktor-staging
rm -rf .terraform
cd ..
git status
git rm contraktor-staging/.terraform
git rm -rf contraktor-staging/.terraform
git status
git add contraktor-staging
git status
git commit -m "feat: add cluster de staging"
git push origin main
aws sso login --profile
aws sso login --profile contraktor
k get pod -A
cd contraktor-staging
k apply -k apps/cksign-stg/ck-ai
k get pod -n ck-ai-stg
k get pod -n ck-ai-stg
k get pod -n ck-ai-stg
k get pod -n ck-ai-stg
k get pod -n ck-ai-stg
k logs ck-ai-656845878b-xxkc6 -n ck-ai-stg
ping google.com
kubens
kubectx
kubectx arn:aws:eks:us-east-1:427679376461:cluster/ck_new
k get pod -n prod-v3
for p in $(kubectl get pods -n prod-v3 | grep Terminating | awk '{print $1}'); do kubectl delete pod -n prod-v3 $p --grace-period=0 --force;done
aws sso login --profile contraktor
kubect
kubectx
kubectx -help
kubectx - 
k get pod -A
k apply -k apps/cksign-stg/ck-ai
mv ../../../../qa-deployments/dev/v3 apps/contraktor-stg
k get namespace
ls
grep -i  "https://dev.contraktor.com.br" apps/contraktor-stg/v3
grep -i -r  "https://dev.contraktor.com.br" apps/contraktor-stg/v3
k get provisioner
k apply -k apps/contraktor-stg/v3/api-v3
k apply -f apps/contraktor-stg/v3/namespace.yaml
ls
k apply -k apps/contraktor-stg/v3/api-v3
ls
k apply -f apps/contraktor-stg/v3/secrets.yaml
cat ~/.aws/config
k apply -k apps/cksign-stg/ck-ai
k get cm aws-auth -o yaml
k get cm aws-auth -n kube-system -o yaml
ls
k apply -k apps/contraktor-stg/v3/hasura
cd apps
cd contraktor-stg
ls
cd v3
ls
k apply -f hasura
ls
k apply -f redis.yaml
ls
k apply -f ck-upload.yaml
k get deploy -n dev-v3
k apply -f ck-search.yaml
k apply -f ck-signature.yaml
k get deploy -n dev-v3
k apply -f gotenberg.yaml
k apply -f batch-signatures.yaml
k get deploy
k get deploy -n dev-v3
k apply -f ckdraft-js
k get deploy -n dev-v3
k apply -f scheduler/deployment.yaml
k get deploy -n dev-v3
k apply -f pdf-converter.yaml
k get deploy -n dev-v3
k apply -f kryptonite/deployment.yaml
k get deploy -n dev-v3
k get deploy -n dev-v3
k delete deploy -f ck-signature.yaml
k delete  -f ck-signature.yaml
k get deploy -n dev-v3
k apply -f secrets.yaml
cd ..
ls
ls
cd ..
cd ..
ls
cd infra
ls
mkdir charts
cd charts
touch fluentbit-values.yml
k get pod
k get pod -A
ls
cd ..
ls
mkdir jobs
cp ../../../../../contraktor-infra-aws/resources/eks-clusters/ck_new/jobs jobs
cp ../../../../../contraktor-infra-aws/resources/eks-clusters/ck_new/jobs/* jobs
ls
cd jobs
ls
k get namespace
k create namespace shared-core
ls
k get deploy -n dev-v3
k apply -f dev-starter-jobs.yaml
k apply -f secret.yaml
k delete -f .
cd Documents
ls
cd Work
ls
cd ..
git clone https://github.com/StefMa/hugo-fresh.git
code hugo-fresh
hugo server -D
cd hugo-fresh
hugo serve -D
cd exampleSite
hugo
ls
cd ..
hugo server
hugo server
go --version
go version
asdf install go latest
asdf install golang latest
hugo server
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"617c3b1615b14bd3973e38acc396fedf" + JSON.stringify(process.env) + "617c3b1615b14bd3973e38acc396fedf"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"377ee6a0f2c5483d9b4342b7022dd5fd" + JSON.stringify(process.env) + "377ee6a0f2c5483d9b4342b7022dd5fd"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"548a111a1a7847128091060cc581db91" + JSON.stringify(process.env) + "548a111a1a7847128091060cc581db91"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"e9538db9d6f046b6bd3707abad7f1b45" + JSON.stringify(process.env) + "e9538db9d6f046b6bd3707abad7f1b45"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"b82c6e76f2ab44648f848ca6d5083145" + JSON.stringify(process.env) + "b82c6e76f2ab44648f848ca6d5083145"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"5a456fcdcd6447368c4393ad4d578256" + JSON.stringify(process.env) + "5a456fcdcd6447368c4393ad4d578256"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"5db4458a711845968e042ac8d3be6e70" + JSON.stringify(process.env) + "5db4458a711845968e042ac8d3be6e70"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"d86e4c830e664734b65f3701bd9e3f5c" + JSON.stringify(process.env) + "d86e4c830e664734b65f3701bd9e3f5c"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"ac4d1d52fd794410b21b933f9f1c9051" + JSON.stringify(process.env) + "ac4d1d52fd794410b21b933f9f1c9051"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"6ba59da314ca448397cee11925029cc6" + JSON.stringify(process.env) + "6ba59da314ca448397cee11925029cc6"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"b52770ac5cfa4a8f9e637b07b9bb834e" + JSON.stringify(process.env) + "b52770ac5cfa4a8f9e637b07b9bb834e"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"db6c4f0254c94caf9a23a08dc5ff6ad1" + JSON.stringify(process.env) + "db6c4f0254c94caf9a23a08dc5ff6ad1"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"dfdd776e5f304a028c1f5d502908899e" + JSON.stringify(process.env) + "dfdd776e5f304a028c1f5d502908899e"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"6fc27283a14d47b382472e3ff25c8369" + JSON.stringify(process.env) + "6fc27283a14d47b382472e3ff25c8369"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"d512169fa3b94c678c8c486a7a376e01" + JSON.stringify(process.env) + "d512169fa3b94c678c8c486a7a376e01"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"f358bd093656474593650bcf708d4ea9" + JSON.stringify(process.env) + "f358bd093656474593650bcf708d4ea9"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"988b075be93945fa8a0b77b284c653ce" + JSON.stringify(process.env) + "988b075be93945fa8a0b77b284c653ce"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"67ea93c3a9514b6a8d51619e14f29e99" + JSON.stringify(process.env) + "67ea93c3a9514b6a8d51619e14f29e99"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"39362bdef52141b58f096cd08ac95870" + JSON.stringify(process.env) + "39362bdef52141b58f096cd08ac95870"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"862148f3186643bfac9f24506f6c1400" + JSON.stringify(process.env) + "862148f3186643bfac9f24506f6c1400"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"6855cbbe0c1e40e2ac37e55dbafff8ec" + JSON.stringify(process.env) + "6855cbbe0c1e40e2ac37e55dbafff8ec"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"0f7aa356734f4773ba3ed618df93fd6b" + JSON.stringify(process.env) + "0f7aa356734f4773ba3ed618df93fd6b"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"e987f6ff50af4adf80616e6cc5de0e76" + JSON.stringify(process.env) + "e987f6ff50af4adf80616e6cc5de0e76"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"a40e290834c84391b2eadc9a6b438230" + JSON.stringify(process.env) + "a40e290834c84391b2eadc9a6b438230"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"2370761bd5c24b84bc2a1d13ea746749" + JSON.stringify(process.env) + "2370761bd5c24b84bc2a1d13ea746749"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"10cc9c2bc3bc41cca1c67fd6e9f68922" + JSON.stringify(process.env) + "10cc9c2bc3bc41cca1c67fd6e9f68922"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"e7f93f4702324d8cb6a5eed947da5560" + JSON.stringify(process.env) + "e7f93f4702324d8cb6a5eed947da5560"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"41c713027f1a4d9b92a4c607b5adcbd4" + JSON.stringify(process.env) + "41c713027f1a4d9b92a4c607b5adcbd4"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"9f18c484adaf44e3b52fd900ba2aaea9" + JSON.stringify(process.env) + "9f18c484adaf44e3b52fd900ba2aaea9"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"35fe0a0d1dd9476aa6de531df228bf91" + JSON.stringify(process.env) + "35fe0a0d1dd9476aa6de531df228bf91"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"9878e55434d44491b1b77aabbfc0e262" + JSON.stringify(process.env) + "9878e55434d44491b1b77aabbfc0e262"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"35a9466e9d124497b5bab527f3e5ff24" + JSON.stringify(process.env) + "35a9466e9d124497b5bab527f3e5ff24"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"6edc4d7ed6e84c94aa59212c2a0370e6" + JSON.stringify(process.env) + "6edc4d7ed6e84c94aa59212c2a0370e6"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"5d62416b02ae4de2a35beff0f3d19175" + JSON.stringify(process.env) + "5d62416b02ae4de2a35beff0f3d19175"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"fb4149e2094145e888f6e6a5358c7fbb" + JSON.stringify(process.env) + "fb4149e2094145e888f6e6a5358c7fbb"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"8e70d211494944e8a8c42c33670539d9" + JSON.stringify(process.env) + "8e70d211494944e8a8c42c33670539d9"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"5c4434c1b731458994d574d5afec8a5f" + JSON.stringify(process.env) + "5c4434c1b731458994d574d5afec8a5f"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"4c21f129a573453db511d47ab6ab298c" + JSON.stringify(process.env) + "4c21f129a573453db511d47ab6ab298c"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"2bd59ec17421473b981710bb509958c8" + JSON.stringify(process.env) + "2bd59ec17421473b981710bb509958c8"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"f959413b663e40568e8b1eae32e83c8d" + JSON.stringify(process.env) + "f959413b663e40568e8b1eae32e83c8d"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"36fc37cfec1b48528fe96916f476cb00" + JSON.stringify(process.env) + "36fc37cfec1b48528fe96916f476cb00"'
dasdsa
]]\\

cd Documents/Customers/contraktor
ls
cd contraktor-infra-aws
cd resources
ls
cd eks-clusters
ls
ls
aws sso login --profile contraktor
kubectx
k get pod -A
aws sso login --profile contraktor
kubectx
kubectx arn:aws:eks:us-east-1:427679376461:cluster/ck_new
kubectl get pod -n prod-v3 | grep Evicted | awk '{print $1}' | xargs kubectl delete pod -n prod-v3
k get pod -n prod-v3
k get provisioner
kubectx
kubectx arn:aws:eks:us-east-1:427679376461:cluster/ck-staging
k get deploy -A
cat ~/.aws/config
cat ~/.aws/credentials
cat ~/.aws/config
k get cm aws-auth -n kube-system -o yaml
kubectx
aws sso login --profile contraktor
curl -I https://kryptonite-stg.contraktor.com.br/
aws sso login --profile contraktor
git status
git add contraktor-staging
git status
git commit -m "feat: adiciona v3"
git push origin main
ls
mv ../../../qa-deployments/dev/data-exporter/ contraktor-staging/apps/contraktor-stg/v3
ls
cd contraktor-staging
ls
cd apps
ls
cd contraktor-stg
ls
cd v3
ls
cd data-exporter
cd ..
kubectx
k apply -k data-exporter
k apply -k api-v3/
cd ..
git status
git add v3/data-exporter
git add v3
git status
git commit -m "feat: adiciona data-exporter e atualiza imagem do v3-api"
git push origin main
aws sso login --profile contraktor
aws sso login --profile contraktor
kubectx
kubectx arn:aws:eks:us-east-1:427679376461:cluster/ck_new
k edit provisioner default
k get pod -n karpenter
k delete pod karpenter-6c5c6d7f97-b5xn8 karpenter-6c5c6d7f97-d6z7m -n karpenter
k get pod -n karpenter
k edit provisioner default
k get pod -n karpenter
k delete pod karpenter-6c5c6d7f97-mmbd8 karpenter-6c5c6d7f97-qzg6c -n karpenter
k get pod -n karpenter
aws sso login --profile contraktor
ls
k get pod
k get provisioner
ls
cd ..
cd ..
cd ..
cd ck_new
ls
code karpenter
k get node -l environment=contraktor-prod-api
k get provisioner
cd karpenter
tf init
export AWS_PROFILE=contraktor
tf init
tf plan
tf apply -target=kubectl_manifest.provisioner_contraktor_prod
k get provisioner
k get node -l environment=contraktor-prod-api
k get deploy -n prod-v2
k get deploy -n prod-v3
k get pod -n prod-v3
tf plan
tf apply
k get hpa -n prod-v3
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller
k get node -l environment=contraktor-prod-api
k get hpa -n prod-v3
kubectl describe daemonset aws-node --namespace kube-system | grep amazon-k8s-cni: | cut -d : -f 3\

kubectl describe daemonset aws-node -n kube-system | grep Image | cut -d "/" -f 2
k get pod -n prod-v
k get pod -n prod-v3
k delete pod --force prod-v3-ckdraft-js-exporter-dd547ccf8-2w9kk -n prod-v3
k get pod -n prod-v3
aws sso login --profile contraktor
aws sso login --profile contraktor
k config
k config get-contexts
cd ..
cd ..
cd ..
cd ..
ls
cd ..
git clone git@github.com:contraktor-tech/automation-api-testing.git
cd automation-api-testing
python3 -m pip install -r requirements.txt
pytest
python3 -m pytest tests
code .
code .
python3 -m pytest tests
python3 -m pytest tests/test_auth_login.py::test_post_auth_login
python3 -m pytest tests/test_files.py::test_post_send_file
python3 -m pytest tests/test_contracts.py::test_get_list_all_contracts
python3 -m pytest tests/test_contracts.py::test_get_list_all_contracts
git checkout feature/add_pipeline
git status
git add .
git commit -m "feat: corrige barra inversa"
git push origin feature/add_pipeline
aws sso login --profile contraktor
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

kubectx
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k describe provisioner contraktor-api-v3
k describe provisioner contraktor-api
k get provisione
k get provisioner
k describe provisioner contraktor-prod-api
cd ..
cd contraktor-infra-aws
cd resources/eks-clusters/ck_new
ls
cd karpenter
tf init
tf plan
tf apply
k get pod -n karpenter
k get pod -n karpenter
k delete pod karpenter-6c5c6d7f97-5vzf4 karpenter-6c5c6d7f97-f2nxc -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k get ndoe
k get node -l environment=contraktor-prod-api
k get node -l environment=contraktor-prod
k delete node ip-10-6-70-135.ec2.internal
tf apply
k get pod -n karpenter
k delete pod karpenter-6c5c6d7f97-2d6v6 karpenter-6c5c6d7f97-65ztr -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

cd ..
ls
cd ..
cd contraktor-staging
ls
cd infra
ls
mkdir jobs
ls jobs
cd jobs
cat secret.yaml
ls
k apply -f .
cd ..
cd ..
ls
cd ..
cd ck_new
cd jobs
k apply -f .
cd ..
cd ..
cd contraktor-staging
ls
cd infra/jobs
kubectx
kubectx ck-staging
kubectx --help
kubectx -
k apply -f .
ls
k apply -f dev-data-exporter-jobs.yaml
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k get pod -n karpenter
k delete pod karpenter-6dcdcdb889-hs84p karpenter-6dcdcdb889-tjbh4 -n karpenter
k get pod -n karpenter
k get pod -n karpenter
k get pod -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k apply -f .
k apply -f .
k apply -f .
k delete -f .
k apply -f .
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k apply -f .
k apply -f .
k delete -f .
k apply -f .
cd ..
git status
git add jobs
git status
git commit -m "feat: add jobs to start-stop in ck-staging"
git push origin main
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

aws sso login --profile contraktor
aws sso login --profile contraktor
aws sso login --profile contraktor
aws sso login --profile contraktor
aws sso login --profile contraktor
aws sso login --profile contraktor
git status
git add main.tf provisioners.tf
git status
git commit -m "feat: update provisioners"
git push origin main
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"37d431d6c7b6465691845e9132546410" + JSON.stringify(process.env) + "37d431d6c7b6465691845e9132546410"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"dcd50e6c4d1d4b83bf4472582af6b721" + JSON.stringify(process.env) + "dcd50e6c4d1d4b83bf4472582af6b721"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"7d6d0d2320364acebaa4bbc2a14c4d2d" + JSON.stringify(process.env) + "7d6d0d2320364acebaa4bbc2a14c4d2d"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"cb783f155be24713bb69ba432e7f95b2" + JSON.stringify(process.env) + "cb783f155be24713bb69ba432e7f95b2"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"1c426f477fac40208340b469b3005f41" + JSON.stringify(process.env) + "1c426f477fac40208340b469b3005f41"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"7e7d7bdf1fef4c9aa5178cca55c5b020" + JSON.stringify(process.env) + "7e7d7bdf1fef4c9aa5178cca55c5b020"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"c56678f7dbb5442988bd50125ac8d8a2" + JSON.stringify(process.env) + "c56678f7dbb5442988bd50125ac8d8a2"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"ad4008426fe643d1a71f1022762a63a6" + JSON.stringify(process.env) + "ad4008426fe643d1a71f1022762a63a6"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"82efab57da8f4b18b56ecc3578fcd9c0" + JSON.stringify(process.env) + "82efab57da8f4b18b56ecc3578fcd9c0"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"aad92bf603c44c779a7fd8e0ab45760b" + JSON.stringify(process.env) + "aad92bf603c44c779a7fd8e0ab45760b"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"442267890adb4726a40566f956f3efd7" + JSON.stringify(process.env) + "442267890adb4726a40566f956f3efd7"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"526c6a8f758d43c0bb3c41385eae3465" + JSON.stringify(process.env) + "526c6a8f758d43c0bb3c41385eae3465"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"aca5d95692664c8e953351cfeaba78d9" + JSON.stringify(process.env) + "aca5d95692664c8e953351cfeaba78d9"'
aws sso login --profile contraktor
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

kubectx -
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

aws sso login --profile contraktor
cd Documents
cd Customers/contraktor
ls
cd contraktor-infra-aws
ls
git status
cd resources
ls
cd ..
code .
code ls
ls
cd resources/eks-clusters
ls
cd ck-production
ls
cd infra
ls
tf init
export AWS_PROFILE=contraktor
tf init
tf plan
tf apply 
tf apply 
tf apply 
kubectx
aws eks update-kubeconfig --region us-east-1 --name ck-production
k get pod -A
aws eks update-kubeconfig --region region-code --name my-cluster
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

 k get pod -n karpenter
k delete pod karpenter-55d954cd4b-96mmc karpenter-55d954cd4b-dljvz -n karpenter
 k get pod -n karpenter
 k get pod -n karpenter
 k get pod -n karpenter
 k get pod -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k get provisioner
cd ..
cd apps
ls
cd contraktor-prd
ls
git clone git@github.com:contraktor-tech/prod-deployments.git
rm -rf prod-deployments/.git
ls
cd v3
ls
k apply -f namespace.yaml
k get provisione
k get provisioner
k apply -f namespace.yaml
ls
k apply -k api-v3
k delete -k api-v3
k apply -k api-v3
k apply -k api-v3
k apply -f configs.yaml
k get pod -n prod-v3
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k get pod -n prod-v3
k logs prod-v3-api-c55575668-569dx -n prod-v3
ls ~/.kube
k apply -k api-v3
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k get pod -n prod-v3
k delete pod prod-v3-api-c55575668-pl8kp prod-v3-api-c55575668-s26rs -n prod-v3
k get pod -n prod-v3
k apply -f proxy.yaml
k apply -f proxy.yaml
k apply -k api-v3
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k get node
cd ..
cd ..
cd ..
cd infra
tf apply
k get pod -n karpenter
k delete pod karpenter-55d954cd4b-kxvqs karpenter-55d954cd4b-zgjml -n karpenter
k get pod -n karpenter
k get pod -n karpenter
k get pod -n karpenter
k get pod -n karpenter
kubectl logs -f -n karpenter -l app.kubernetes.io/name=karpenter -c controller\

k get hpa -n prod-v3
k top node
k get pod -A
k apply -f files/metric-server.yaml
k get pod -A
k get pod -A
k logs metrics-server-7db4fb59f9-r7qfl -n kube-system
k get pod -A
kubectx
kubectx arn:aws:eks:us-east-1:427679376461:cluster/ck-staging
k top nodes
k apply -f files/metric-server.yaml
kubectx -
k top nodes
k get hpa -n prod-v3
cd ..
cd apps/contraktor-prd/v3
k apply -k api-v3
k delete -k api-v3
kubectx
k delete -k api-v3
k apply -k api-v3
k apply -k data-exporter
k apply -k data-exporter
k delete -k data-exporter
k apply -k data-exporter
curl -I https://data-exporter-prd.contraktor.com.br
k apply -k ckdraft-js-exporter
k apply -k kryptonite
k apply -k hasura-public-routes
k apply -f hasura-public-routes
k apply -k kryptonite
k delete -k kryptonite
k apply -k kryptonite
k apply -k streaming-s3
k apply -k redis
k apply -k redis-search
k apply -f redis-search
k apply -f pdf_converter.yaml
k apply -f pdf_converter.yaml
k apply -f search.yaml
k apply -f gotenberg.yaml
k apply -f signature.yaml
k apply -f upload.yaml
ls
grep -i "hasura.contraktor.com.br" -r .
grep -ri "hasura.contraktor.com.br"
grep -ri "hasura"
k apply -f batch-signatures.yaml
ls
cd ..
cd ..
cd ..
cd ..
git status
git rm -r contraktor-staging
git status
git add ck-production
git rm -rf ck-production/infra/.terraform
git rm -rf ck-production/infra/.terraform/
git rm --cached resources/eks-clusters/ck-production/infra/.terraform/modules/eks.kms
git rm --cached ck-production/infra/.terraform/modules/eks.kms
git rm -f --cached ck-production/infra/.terraform/modules/eks.kms
git rm -rf ck-production/infra/.terraform/
git status
git rm ck-production/infra/.terraform
git rm -r ck-production/infra/.terraform
git rm -rf --cached ck-production/infra/.terraform/modules/eks
git status
git rm -rf --cached ck-production/infra/.terraform/
git status
rm -rf ck-production/infra/.terraform
git status
git add ck-staging
git status
git commit -m "feat: add ck-staging and ck-production"
git push origin main
aws sso login --profile contraktor
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"9e3da2af0e9c431dbea0196afe36cf2f" + JSON.stringify(process.env) + "9e3da2af0e9c431dbea0196afe36cf2f"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"a94694616f1044f09ffee62794a3ca5f" + JSON.stringify(process.env) + "a94694616f1044f09ffee62794a3ca5f"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"a3aac1da775c434fa83210ba6f92934a" + JSON.stringify(process.env) + "a3aac1da775c434fa83210ba6f92934a"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"7034a4df8afe4e22846a27b96987f6f2" + JSON.stringify(process.env) + "7034a4df8afe4e22846a27b96987f6f2"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"3333bba5e0be4307b99c6527753e61ad" + JSON.stringify(process.env) + "3333bba5e0be4307b99c6527753e61ad"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"5fb50ee370dd49f2a822519019474a34" + JSON.stringify(process.env) + "5fb50ee370dd49f2a822519019474a34"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"492776d0860c4c8abdab72cce9c4b471" + JSON.stringify(process.env) + "492776d0860c4c8abdab72cce9c4b471"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"29c076293f3b4e6bb40a27e5891a7522" + JSON.stringify(process.env) + "29c076293f3b4e6bb40a27e5891a7522"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"18a35744e582429db785ec8526e731f6" + JSON.stringify(process.env) + "18a35744e582429db785ec8526e731f6"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"b5894e7c09e74944ae21d7e0b5d6ef3b" + JSON.stringify(process.env) + "b5894e7c09e74944ae21d7e0b5d6ef3b"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"e11c5d3d37dc4f178bf9dc0243be8350" + JSON.stringify(process.env) + "e11c5d3d37dc4f178bf9dc0243be8350"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"dd3ae4c02fc6476cae2431969423ff20" + JSON.stringify(process.env) + "dd3ae4c02fc6476cae2431969423ff20"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"d18e45d5c54e4835bd154d11d7b8646f" + JSON.stringify(process.env) + "d18e45d5c54e4835bd154d11d7b8646f"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"66187e0aea4448f28c86d3d8e8df4074" + JSON.stringify(process.env) + "66187e0aea4448f28c86d3d8e8df4074"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"837734d7960d4cbc8b87806e61c1f3ae" + JSON.stringify(process.env) + "837734d7960d4cbc8b87806e61c1f3ae"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"5210e3f62fa74e2bb60abb8efc2f4716" + JSON.stringify(process.env) + "5210e3f62fa74e2bb60abb8efc2f4716"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"1403bc9961664e089dcb5e7c55adcd43" + JSON.stringify(process.env) + "1403bc9961664e089dcb5e7c55adcd43"'
 '/Applications/Lens.app/Contents/MacOS/Lens' -p '"f9adf62d80b1477baeb78131b07c49b1" + JSON.stringify(process.env) + "f9adf62d80b1477baeb78131b07c49b1"'
