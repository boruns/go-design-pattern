# 适配器模式

## 定义

    主要用来做适配的，将不兼容的接口转换为可兼容的接口，让原本由于接口不兼容不能一起工作的类可以一起工作。

## 实现方式

1. 类适配器，继承实现
2. 对象适配器，组合实现。

## 如何选择

1. 需要适配的接口不多，两种都可以。
2. 需要适配的接口多，需要大量修改选择对象适配器，不需要大量修改，使用类适配器。


## 应用场景

1. 封装有缺陷的接口设计
2. 统一多个类的接口设计
3. 替换依赖的外部系统
4. 兼容老版本接口
5. 适配不同的格式数据

## 区别

1. 代理模式：代理模式在不改变原始类接口的情况下，为原始类定义一个代理类，主要目的是控制访问，而不是增强功能，这是它与装饰器模式最大的不同。
2. 桥接模式：桥接模式的目的是将接口部分与实现部分分离，从而使它们相对容易、也相对独立的加以改变。
3. 装饰器模式：装饰器在不改变原始类的情况下，对原始类的功能进行增强。
4. 适配器模式：适配器模式是一种事后补救的策略。适配器提供和原始类不同的接口，而代理模式、装饰器模式提供的都是和原始类相同的接口。
