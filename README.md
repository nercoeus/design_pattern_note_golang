# 什么是设计模式
&emsp;&emsp;[设计模式](https://zh.wikipedia.org/wiki/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F_(%E8%AE%A1%E7%AE%97%E6%9C%BA))（design pattern）是对软件设计中普遍存在（反复出现）的各种问题，所提出的解决方案.  
&emsp;&emsp;设计模式并不是教你如何编写代码，而是在描述各种不同的情况下，要怎样解决问题的方式  
&emsp;&emsp;每一个模式描述了一个在我们周围不断重复发生的问题，以及该问题的核心解决办法   -----Christopher Alexander  

# 源码

博客中所有源码见 /src

# 设计模式的种类
本书中给出了 23 种设计模式：

name   | describe 
----- | --- 
[Abstract Factory](https://nercoeus.github.io/2019/10/31/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E6%8A%BD%E8%B1%A1%E5%B7%A5%E5%8E%82/) | 提供一个创建一系列相关或相互依赖对象的接口，而无需制定具体的类
[Adapter Pattern](https://nercoeus.github.io/2019/11/02/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E9%80%82%E9%85%8D%E5%99%A8%E6%A8%A1%E5%BC%8F/)| 将一个类的接口转换成客户希望的另一个接口。使得原本接口不兼容的类可以一起工作
[Bridge Pattern](https://nercoeus.github.io/2019/11/04/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E6%A1%A5%E6%8E%A5%E6%A8%A1%E5%BC%8F/) | 将一个抽象与实现解耦，以便两者可以独立的变化
[Builder Pattern](https://nercoeus.github.io/2019/10/31/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E6%8A%BD%E8%B1%A1%E5%B7%A5%E5%8E%82/) | 将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示
[Chain-of-responsibility pattern](https://nercoeus.github.io/2019/11/05/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E8%81%8C%E8%B4%A3%E9%93%BE%E6%A8%A1%E5%BC%8F/) | 为解除请求的发送者和接收者之间耦合，而使多个对象都有机会处理这个请求。将这些对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理它
[Command pattern](https://nercoeus.github.io/2019/11/05/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E5%91%BD%E4%BB%A4%E6%A8%A1%E5%BC%8F/) | 将一个请求封装为一个对象，从而使你可用不同的请求对客户进行参数化；对请求排队或记录请求日志，以及支持可取消的操作
[Composite pattern](https://nercoeus.github.io/2019/11/04/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E7%BB%84%E5%90%88%E6%A8%A1%E5%BC%8F/) | 把多个对象组成树状结构来表示局部与整体，这样用户可以一样的对待单个对象和对象的组合
[Decorator pattern](https://nercoeus.github.io/2019/11/04/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E8%A3%85%E9%A5%B0%E6%A8%A1%E5%BC%8F/) | 向某个对象动态地添加更多的功能。修饰模式是除类继承外另一种扩展功能的方法
[Façade pattern](https://nercoeus.github.io/2019/11/04/设计模式-外观模式/) | 为子系统中的一组接口提供一个一致的界面， 外观模式定义了一个高层接口，这个接口使得这一子系统更加容易使用
[Factory Method pattern](https://nercoeus.github.io/2019/10/31/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E5%B7%A5%E5%8E%82%E6%96%B9%E6%B3%95/) | 定义一个接口用于创建对象，但是让子类决定初始化哪个类。工厂方法把一个类的初始化下放到子类
[Flyweight pattern](https://nercoeus.github.io/2019/11/05/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E4%BA%AB%E5%85%83%E6%A8%A1%E5%BC%8F/) | 通过共享以便有效的支持大量小颗粒对象
[Interpreter pattern](https://nercoeus.github.io/2019/11/06/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E8%A7%A3%E9%87%8A%E5%99%A8%E6%A8%A1%E5%BC%8F/) | 给定一个语言, 定义它的文法的一种表示，并定义一个解释器, 该解释器使用该表示来解释语言中的句子
[Iterator pattern](https://nercoeus.github.io/2019/11/06/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E8%BF%AD%E4%BB%A3%E5%99%A8%E6%A8%A1%E5%BC%8F/) | 提供一种方法顺序访问一个聚合对象中各个元素, 而又不需暴露该对象的内部表示
[Mediator pattern](https://nercoeus.github.io/2019/11/13/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E4%B8%AD%E4%BB%8B%E8%80%85%E6%A8%A1%E5%BC%8F/) | 包装了一系列对象相互作用的方式，使得这些对象不必相互明显作用，从而使它们可以松散偶合。当某些对象之间的作用发生改变时，不会立即影响其他的一些对象之间的作用，保证这些作用可以彼此独立的变化
[Memento pattern](https://nercoeus.github.io/2019/11/13/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E5%A4%87%E5%BF%98%E5%BD%95%E6%A8%A1%E5%BC%8F/) | 回忆对象是一个用来存储另外一个对象内部状态的快照的对象。回忆模式的用意是在不破坏封装的条件下，将一个对象的状态捉住，并外部化，存储起来，从而可以在将来合适的时候把这个对象还原到存储起来的状态
[Observer pattern](https://nercoeus.github.io/2019/11/14/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E8%A7%82%E5%AF%9F%E8%80%85%E6%A8%A1%E5%BC%8F/) | 在对象间定义一个一对多的联系性，由此当一个对象改变了状态，所有其他相关的对象会被通知并且自动刷新
[Prototype pattern](https://nercoeus.github.io/2019/11/02/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E5%8E%9F%E5%9E%8B%E6%A8%A1%E5%BC%8F/) | 用原型实例指定创建对象的种类，并且通过拷贝这些原型,创建新的对象
[Proxy pattern](https://nercoeus.github.io/2019/11/05/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E4%BB%A3%E7%90%86%E6%A8%A1%E5%BC%8F/) | 为其他对象提供一个代理以控制对这个对象的访问
[Singleton pattern](https://nercoeus.github.io/2019/11/02/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E5%8D%95%E4%BE%8B%E6%A8%A1%E5%BC%8F/) | 确保一个类只有一个实例，并提供对该实例的全局访问
[State pattern](https://nercoeus.github.io/2019/11/14/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E7%8A%B6%E6%80%81%E6%A8%A1%E5%BC%8F/) | 让一个对象在其内部状态改变的时候，其行为也随之改变。状态模式需要对每一个系统可能获取的状态创立一个状态类的子类。当系统的状态变化时，系统便改变所选的子类
[Strategy pattern](https://nercoeus.github.io/2019/11/14/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E7%AD%96%E7%95%A5%E6%A8%A1%E5%BC%8F/) | 定义一个算法的系列，将其各个分装，并且使他们有交互性。策略模式使得算法在用户使用的时候能独立的改变
[Template method pattern](https://nercoeus.github.io/2019/11/14/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E6%A8%A1%E6%9D%BF%E6%96%B9%E6%B3%95%E6%A8%A1%E5%BC%8F/) | 模板方法模式准备一个抽象类，将部分逻辑以具体方法及具体构造子类的形式实现，然后声明一些抽象方法来迫使子类实现剩余的逻辑。不同的子类可以以不同的方式实现这些抽象方法，从而对剩余的逻辑有不同的实现。先构建一个顶级逻辑框架，而将逻辑的细节留给具体的子类去实现
[Visitor](https://nercoeus.github.io/2019/11/14/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F-%E8%AE%BF%E9%97%AE%E8%80%85%E6%A8%A1%E5%BC%8F/) | 封装一些施加于某种数据结构元素之上的操作。一旦这些操作需要修改，接受这个操作的数据结构可以保持不变。访问者模式适用于数据结构相对未定的系统，它把数据结构和作用于结构上的操作之间的耦合解脱开，使得操作集合可以相对自由的演化

# 设计模式的使用方法
&emsp;&emsp;书第一章的后面推荐了一种使用设计模式的方法：
1. **大致浏览一遍模式**：特别注意它的适用性部分和效果部分，确定适合你需要解决的问题
2. **研究其结构部分，参与者部分和协作部分**：雀稗你理解这个设计模式的类和对象以及它们之间是怎么联系的
3. **看代码实例，来掌握它的实际使用**：有助于你实现它
4. **选择模式参与者的名字，使它们在应用的上下文之间有意义**：不要直接使用设计模式的名字，因为它过于抽象，要根据你自己的工程上下文设计有意义的名字，可以和设计模式的名字进行组合
5. **定义类**：根据设计模式的类和对象结构，将其融入到你自己的工程中
6. **定义模式中专用于应用的操作名称**：根据上一步定义的类，来定义其对应的操作
7. **实现执行模式中责任和写作的操作**：实现部分提供线索指导你进行实现

# 最后
&emsp;&emsp;设计模式每一个单独来看均不是十分困难,难点在于要学会灵活运用,在日常编写代码的过程中进行使用,理解这些设计模式所善于解决的问题和所带来的弊端.

# 参考
1. 《设计模式》
2. [Dive into design patterns](https://refactoring.guru/design-patterns)
3. [参考2的中文版，还在完善中](https://refactoringguru.cn/design-patterns)
4. [设计模式 Golang 实现](https://github.com/senghoo/golang-design-pattern)