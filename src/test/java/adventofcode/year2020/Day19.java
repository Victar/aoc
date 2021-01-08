package adventofcode.year2020;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;
import lombok.ToString;

public class Day19 extends BaseTest {

	@Test public void singleCheck() {
	}

	@Test
	public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day19/input.txt");
		ArrayList<Rule> rulesList = new ArrayList<>();
		ArrayList<String> messageList = new ArrayList<>();
		fillData(data, rulesList, messageList);
		rulesList.sort(Comparator.comparing(a -> a.getNum()));
		Rule rule0 = rulesList.stream().filter(l -> l.getNum() == 0).findFirst().orElse(null);
		String regExp0 = rule0.toRegExpString();
		//		System.out.println(regExp0);
		int count = 0;
		for (String message : messageList){
			if (message.matches(regExp0)){
				count++;
			}
		}
		System.out.println(count);
	}

	@Test
	public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day19/input_replace.txt");
		ArrayList<Rule> rulesList = new ArrayList<>();
		ArrayList<String> messageList = new ArrayList<>();
		fillData(data, rulesList, messageList);
		rulesList.sort(Comparator.comparing(a -> a.getNum()));
		Rule rule0 = rulesList.stream().filter(l -> l.getNum() == 0).findFirst().orElse(null);
		String regExp0 = rule0.toRegExpString();
//		System.out.println(regExp0);
		int count = 0;
		for (String message : messageList){
			if (message.matches(regExp0)){
				count++;
			}
		}
		System.out.println(count);
	}

	public void fillData(final ArrayList<String> data, final ArrayList<Rule> rulesList, final ArrayList<String> messageList){
		for (String input : data) {
			if (input.contains(":")) {
				//parse rule
				String[] inputArray = StringUtils.split(input, ":\\|");
				int ruleNumber = Integer.parseInt(inputArray[0]);
				Rule currentRule = rulesList.stream().filter(l -> l.getNum()==ruleNumber).findFirst().orElse(null);
				if (currentRule == null) {
					currentRule = new Rule(ruleNumber);
					rulesList.add(currentRule);
				}
				String input2 = inputArray[1];
				if (input2.contains("a")){
					currentRule.setC("a");
				}else if(input2.contains("b")){
					currentRule.setC("b");
				}else {
					String[] childLeft = inputArray[1].split(" ");
					for (int i = 0; i < childLeft.length; i++) {
						if (StringUtils.isNotEmpty(childLeft[i])) {
							int childRuleNumber = Integer.parseInt(childLeft[i].trim());
							Rule childRule = rulesList.stream().filter(l -> l.getNum() == childRuleNumber).findFirst().orElse(null);
							if (childRule == null) {
								childRule = new Rule(childRuleNumber);
								rulesList.add(childRule);
							}
							currentRule.addLeft(childRule);
						}
					}
				}
				if (inputArray.length ==3){
					String[]  childRight = inputArray[2].split(" ");
					for (int i = 0; i < childRight.length; i++) {
						if (StringUtils.isNotEmpty(childRight[i])) {

							int childRuleNumber = Integer.parseInt(childRight[i].trim());
							Rule childRule =
									rulesList.stream().filter(l -> l.getNum() == childRuleNumber).findFirst().orElse(null);
							if (childRule == null) {
								childRule = new Rule(childRuleNumber);
								rulesList.add(childRule);
							}
							currentRule.addRight(childRule);
						}
					}
				}
			}else{
				if (StringUtils.isNotEmpty(input)){
					messageList.add(input);
				}
			}
		}
	}




	@Data
	@ToString
	class Rule{
		int num;
		List<Rule> left = new ArrayList<>();
		List<Rule> right = new ArrayList<>();
		String c = null;

		int  goldVisitedTime = 0;

		public String toRegExpString() {
			if (StringUtils.isNotEmpty(c)){
				return c;
			}
			StringBuilder sb = new StringBuilder("(");
			for (Rule rule : left) {
				if (num!=rule.num || goldVisitedTime<10) {
					if (num==rule.num) {
						goldVisitedTime++;
					}
					sb.append("(" + rule.toRegExpString() + ")");
				}
			}
			sb.append(")");
			if (right.size() > 0){
				sb.append("|(");
				for (Rule rule : right) {
					if (num!=rule.num || goldVisitedTime<10) {
						if (num==rule.num) {
							goldVisitedTime++;
						}
						sb.append("(" + rule.toRegExpString() + ")");
					}
				}
				sb.append(")");
			}
			return sb.toString();
		}


		public Rule(final int num) {
			this.num = num;
		}
		public void addLeft(Rule rule) {
			left.add(rule);
		}
		public void addRight(Rule rule) {
			right.add(rule);
		}

		public void setC(String c){
			this.c = c;
		}

		@Override public String toString() {
			return "Rule{" + "num=" + num + ", left=" + left.size() + ", right=" + right.size() + ", c='" + c + '\'' + '}';
		}
	}
}

