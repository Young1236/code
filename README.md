该项目是一个简单的小的登录系统，用gin框架建成，通过数个接口实现功能
register接口：实现注册功能。通过gorm创建一张User数据表，包括userId,name,password字段。使用时通过query获取数据，key为name和password.如果该数据表中已经存在输入的name，那么接口会显示注册失败，如果没有相同name,则注册成功。
login接口：实现登录功能。同样使用query获取数据，输入name和password，将输入的数据和User数据表的数据进行比对，如果不存在name，则显示无用户，如果输入的password与数据表中的password不匹配，则显示密码错误
problem接口：实现提问功能。创建一张Problems数据表，包括problemId,problem,name字段。使用时通过query获取数据，problem用于储存问题内容，name用于储存提问者的用户名，problemId作为每个问题独有的值,问题创建完成后显示问题序号和问题内容。
answer接口：实现评论功能。创建一张Answers数据表，包括answerId,problemId,answer,answerer字段。使用时通过query获取数据，answer储存评论内容，answerer用于储存评论用户名字，problemId储存输入的问题序号，answerId作为每条评论独有的值，输入数据后显示评论序号和评论内容。
findProblem接口：实现查找问题功能。通过query,获取提问者的用户名，以此查找问题内容，显示问题序号和问题内容。
findAnswer接口：实现查找评论功能。通过query,获取评论者的用户名，以此查找评论，显示问题的序号以及评论的内容。
updateProblem接口：实现更新问题功能。通过query,获取问题的Id,以此进行问题内容更新，显示问题序号和问题内容。
updateAnswer接口：实现更新评论功能。通过query,获取评论的Id,以此进行评论内容更新，显示评论序号和评论内容。
deleteProblem接口：实现删除问题的功能。通过query,获取问题Id,如果problems数据表中没有该问题，则显示无该问题，如果正常，则显示删除成功。
deleteAnswer接口：实现删除评论的功能。通过query,获取评论Id,如果Answers数据表中没有该评论，则显示无该评论，如果正常，则显示删除成功。